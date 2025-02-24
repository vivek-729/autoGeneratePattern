package node

type Node struct {
	Lhs            *Node          `json:"lhs,omitempty"`
	Operator       string         `json:"operator,omitempty"`
	Rhs            *Node          `json:"rhs,omitempty"`
	TokenIndex     int64          `json:"tokenIndex,omitempty"`
	Instrument     string         `json:"instrument,omitempty"`
	TimeFrame      string         `json:"timeFrame,omitempty"`
	Params         map[string]any `json:"params,omitempty"`
	Value          float64        `json:"value,omitempty"`
	ValueString    string         `json:"valueString,omitempty"`
	ValueTimestamp string         `json:"valueTimestamp,omitempty"`
}

// improvement: can make short form of instrument as O, H, L, C and convert them when creating node for better readability of formula
func OHLC(instrument string, offset int64) Node {
	node := Node{
		Instrument: instrument,
	}

	if offset > 0 {
		params := map[string]any{
			"offset": offset,
		}
		node.Params = params
	}

	return node
}

func NV(num float64) Node {
	node := Node{
		Value: num,
	}
	return node
}

// average instrument creator
func AvgOHLC(instrumentName string, OHLCSource string, length int, smoothingLength int) Node {
	if !(instrumentName == "ma" || instrumentName == "ema" || instrumentName == "rma" || instrumentName == "sma") {
		panic("Invalid average string.")
	}

	node := Node{
		Instrument: instrumentName,
		Params: map[string]any{
			"length": length,
			"source": OHLCSource,
		},
	}

	if smoothingLength > 0 {
		node.Params["smoothingLength"] = smoothingLength
	}

	return node
}

// https://help.tc2000.com/m/69445/l/798550-doji-candle
// 20 * (ABS(O - C)) <= H - L
// 20 * (O(-)C) <= (H - L)
var DojiFormula = []any{NV(20), "*", "(", OHLC("open", 0), "(-)", OHLC("close", 0), ")", "<=", "(", OHLC("high", 0), "-", OHLC("low", 0), ")"}

// https://help.tc2000.com/m/69445/l/800589-bullish-candlestick-patterns-formulas-table
// O1 > C1 AND 10 * (C - O) >= 7 * (H - L) AND C > O1 AND C1 > O AND 10 * (H - L) >= 12 * (AVGH10 - AVGL10)

// ( O1 > C1 ) and
// ( 10 * (C - O) >= 7 * (H - L) ) and
// ( C > O1 ) and
// ( C1 > O ) and
// ( 10 * (H - L) >= 12 * (AVGH10 - AVGL10) )
var BullishEngulfing = []any{"(", OHLC("open", 1), ">", OHLC("close", 1), ")", "and",
	"(", NV(10), "*", "(", OHLC("close", 0), "-", OHLC("open", 0), ")", ">=", NV(7), "*", "(", OHLC("high", 0), "-", OHLC("low", 0), ")", ")", "and",
	"(", OHLC("close", 0), ">", OHLC("open", 1), ")", "and",
	"(", OHLC("close", 1), ">", OHLC("open", 0), ")", "and",
	"(", NV(10), "*", "(", OHLC("high", 0), "-", OHLC("low", 0), ")", ">=", NV(12), "*", "(", AvgOHLC("ma", "high", 10, 1), "-", AvgOHLC("ma", "low", 10, 1), ")", ")",
}
