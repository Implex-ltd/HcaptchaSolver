package hcaptcha

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/Implex-ltd/hcsolver/internal/utils"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var boxes = map[int]Box{
	0: {Start: Point{131, 282, 0}, End: Point{177, 310, 0}},
	1: {Start: Point{250, 274, 0}, End: Point{313, 318, 0}},
	2: {Start: Point{390, 274, 0}, End: Point{438, 324, 0}},
	3: {Start: Point{122, 408, 0}, End: Point{187, 456, 0}},
	4: {Start: Point{250, 400, 0}, End: Point{314, 451, 0}},
	5: {Start: Point{386, 400, 0}, End: Point{448, 466, 0}},
	6: {Start: Point{124, 530, 0}, End: Point{188, 584, 0}},
	7: {Start: Point{250, 539, 0}, End: Point{313, 588, 0}},
	8: {Start: Point{387, 537, 0}, End: Point{446, 579, 0}},
}

func RandomPointInBox(box Box) Point {
	const minDiff = 1
	xDiff := box.End.X - box.Start.X
	if xDiff < 0 {
		xDiff = -xDiff
	}
	yDiff := box.End.Y - box.Start.Y
	if yDiff < 0 {
		yDiff = -yDiff
	}
	return Point{
		X: box.Start.X + rand.Int63n(xDiff+minDiff),
		Y: box.Start.Y + rand.Int63n(yDiff+minDiff),
		T: time.Now().UnixNano() / 1e6,
	}
}

func calculateBezier(WnTime float64, start, ctrl1, ctrl2, end int64) int64 {
	u := 1.0 - WnTime
	tt := WnTime * WnTime
	uu := u * u
	uuu := uu * u
	ttt := tt * WnTime

	res := uuu*float64(start) + 3*uu*WnTime*float64(ctrl1) + 3*u*tt*float64(ctrl2) + ttt*float64(end)
	return int64(res)
}

func GenerateMousePath(start, end Point, numPoints int) []Point {
	const minDiff = 1

	ctrl1 := Point{
		X: start.X + max(rand.Int63n(max((end.X-start.X)/2, minDiff)), minDiff),
		Y: start.Y + max(rand.Int63n(max((end.Y-start.Y)/2, minDiff)), minDiff),
		T: start.T + (end.T-start.T)/4,
	}
	ctrl2 := Point{
		X: ctrl1.X + max(rand.Int63n(max((end.X-ctrl1.X)/2, minDiff)), minDiff),
		Y: ctrl1.Y + max(rand.Int63n(max((end.Y-ctrl1.Y)/2, minDiff)), minDiff),
		T: ctrl1.T + (end.T-ctrl1.T)/2,
	}

	var path []Point
	for i := 0; i < numPoints; i++ {
		WnTime := float64(i) / float64(numPoints-1)

		path = append(path, Point{
			X: calculateBezier(WnTime, start.X, ctrl1.X, ctrl2.X, end.X),
			Y: calculateBezier(WnTime, start.Y, ctrl1.Y, ctrl2.Y, end.Y),
			T: int64((1.0-WnTime)*float64(start.T) + WnTime*float64(end.T)),
		})
	}

	return path
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func PlotPoints(points [][]int64) {
	var path []Point

	for _, p := range points {
		path = append(path, Point{
			X: p[0],
			Y: p[1],
			T: p[2],
		})
	}

	p := plot.New()

	p.Title.Text = "Mouse Path"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := make(plotter.XYs, len(path))
	for i, point := range path {
		pts[i].X = float64(point.X)
		pts[i].Y = float64(point.Y)
	}

	minX, maxX, minY, maxY := pts[0].X, pts[0].X, pts[0].Y, pts[0].Y
	for _, point := range pts[1:] {
		if point.X < minX {
			minX = point.X
		}
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y < minY {
			minY = point.Y
		}
		if point.Y > maxY {
			maxY = point.Y
		}
	}

	p.X.Min = minX - 1
	p.X.Max = maxX + 1
	p.Y.Min = minY - 1
	p.Y.Max = maxY + 1

	s, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}

	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	l, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	p.Add(s, l)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("./.tmp/%s.png", utils.RandomString(10))); err != nil {
		panic(err)
	}

	fmt.Println("Plot saved to mouse_path.png")
}

func Point2path(p []Point) [][]int64 {
	convertedPath := make([][]int64, len(p))

	for i, point := range p {
		convertedPath[i] = []int64{point.X, point.Y, point.T}
	}

	return convertedPath
}

func Click(boxesToClick []int, startTime, duration int64, curveAmount int) [][]int64 {
	var path []Point
	timeIncrement := duration / int64(len(boxesToClick))

	for i, boxNum := range boxesToClick {
		box := boxes[boxNum]
		targetPoint := RandomPointInBox(box)
		targetPoint.T = startTime + timeIncrement*int64(i)

		if i > 0 {
			intermediatePath := GenerateMousePath(path[len(path)-1], targetPoint, curveAmount)
			timeDiff := targetPoint.T - path[len(path)-1].T

			for j, point := range intermediatePath {
				point.T = path[len(path)-1].T + timeDiff*int64(j)/int64(len(intermediatePath))
				intermediatePath[j] = point
			}
			path = append(path, intermediatePath...)
		}

		path = append(path, targetPoint)
	}

	return Point2path(path)
}

func addTime(st int64, toAdd time.Duration) int64 {
	return st + toAdd.Milliseconds()
}

func GetRandomBox() Box {
	boxIDs := make([]int, 0, len(boxes))
	for boxID := range boxes {
		boxIDs = append(boxIDs, boxID)
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(boxIDs))
	randomBoxID := boxIDs[randomIndex]

	return boxes[randomBoxID]
}

func calculateMeanPeriod(events [][]int64) float64 {
	var timeDifferences []int64
	for i := 0; i < len(events)-1; i++ {
		timeDifference := events[i+1][2] - events[i][2]
		timeDifferences = append(timeDifferences, timeDifference)
	}
	var sum int64 = 0
	for _, timeDifference := range timeDifferences {
		sum += timeDifference
	}
	meanPeriod := float64(sum) / float64(len(timeDifferences))
	return meanPeriod
}

// Todo: implement
// generateRandomBrowserSize generates a random browser width and height. minSize and maxSize specify the range of possible values.
func generateRandomBrowserSize(minSize, maxSize int) (width, height int64) {
	aspectRatios := []float64{16.0 / 9.0, 16.0 / 10.0, 4.0 / 3.0}

	rand.Shuffle(len(aspectRatios), func(i, j int) {
		aspectRatios[i], aspectRatios[j] = aspectRatios[j], aspectRatios[i]
	})

	aspectRatio := aspectRatios[0]

	width = int64(rand.Intn(maxSize-minSize+1) + minSize)
	height = int64(float64(width) / aspectRatio)

	return width, height
}

func genBoxToClick(answers map[string]string) []int {
	var num = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var items []int

	rand.Shuffle(len(num), func(i, j int) {
		num[i], num[j] = num[j], num[i]
	})

	for i := 0; i < len(answers) && i < len(num); i++ {
		items = append(items, num[i])
	}

	return items
}

/*
* Todo: add right mouse moovements side, if box is lower/higher edit path to add right box
 */
func (c *Hcap) NewMotionData(m *Motion) string {
	st := time.Now().UnixNano() / 1e6
	duration := int64(utils.RandomNumber(100, 500))

	if !m.IsCheck {
		m.Answers = map[string]string{"x": "x", "y": "y", "z": "z", "d": "z", "a": "z", "i": "z"}
	}

	for i := 0; i < utils.RandomNumber(1, 5); i++ {
		m.Answers[utils.RandomString(5)] = "x"
	}

	toClick := genBoxToClick(m.Answers)

	// Generate all mouse curves
	CaptchaPath := Click(toClick, st, duration, utils.RandomNumber(3, 6))
	MdPath := Click([]int{utils.RandomNumber(0, 8), utils.RandomNumber(0, 8), utils.RandomNumber(0, 8)}, st, duration, utils.RandomNumber(8, 16))
	MuPath := Click([]int{utils.RandomNumber(0, 8), utils.RandomNumber(0, 8), utils.RandomNumber(0, 8)}, st, duration, utils.RandomNumber(3, 10))
	MmPath := Click([]int{utils.RandomNumber(0, 8), utils.RandomNumber(0, 8), utils.RandomNumber(0, 8)}, st, duration, utils.RandomNumber(10, 20))
	//WnTime := time.Duration(utils.RandomNumber(20, 35)) * time.Millisecond

	//PlotPoints(CaptchaPath)

	topLevel := TopLevel{
		St: st,
		Sc: Sc{
			AvailWidth:  int64(c.ChallengeFingerprint.Profile.Screen.AvailWidth),
			AvailHeight: int64(c.ChallengeFingerprint.Profile.Screen.AvailHeight),
			Width:       int64(c.ChallengeFingerprint.Profile.Screen.Width),
			Height:      int64(c.ChallengeFingerprint.Profile.Screen.Height),
			ColorDepth:  int64(c.ChallengeFingerprint.Profile.Screen.ColorDepth),
			PixelDepth:  int64(c.ChallengeFingerprint.Profile.Screen.PixelDepth),
			AvailLeft:   int64(c.Fingerprint.Screen.AvailLeft),
			AvailTop:    int64(c.Fingerprint.Screen.AvailTop),
			Onchange:    nil,
			IsExtended:  true,
		},
		Nv: Nv{
			HardwareConcurrency: int64(c.ChallengeFingerprint.Profile.Misc.HardwareConcurrency),
			DeviceMemory:        int64(c.ChallengeFingerprint.Profile.Misc.DeviceMemory),
			Webdriver:           false,
			MaxTouchPoints:      c.ChallengeFingerprint.Profile.Navigator.MaxTouchPoints,
			CookieEnabled:       true,
			AppCodeName:         c.Fingerprint.Navigator.AppCodeName,
			AppName:             c.Fingerprint.Navigator.AppName,
			AppVersion:          c.ChallengeFingerprint.Profile.Misc.AppVersion,
			Platform:            c.ChallengeFingerprint.Profile.Navigator.Platform,
			Product:             c.Fingerprint.Navigator.Product,
			ProductSub:          c.Fingerprint.Navigator.ProductSub,
			UserAgent:           c.ChallengeFingerprint.UserAgent,
			Vendor:              c.Fingerprint.Navigator.Vendor,
			VendorSub:           c.Fingerprint.Navigator.VendorSub,
			Language:            c.ChallengeFingerprint.Profile.Navigator.Language,
			Languages:           c.ChallengeFingerprint.Profile.Navigator.Languages,
			OnLine:              true,
			PDFViewerEnabled:    c.ChallengeFingerprint.Profile.Misc.PDFViewerEnabled,
			DoNotTrack:          c.Fingerprint.Navigator.DoNotTrack,
			Plugins:             []string{"internal-pdf-viewer", "internal-pdf-viewer", "internal-pdf-viewer", "internal-pdf-viewer", "internal-pdf-viewer"},
			UserAgentData: UserAgentData{
				Brands: []Brand{
					{
						Brand:   "Not/A)Brand",
						Version: "8",
					},
					{
						Brand:   "Google Chrome",
						Version: c.Http.BaseHeader.UaInfo.UaVersion,
					},
					{
						Brand:   "Chromium",
						Version: c.Http.BaseHeader.UaInfo.UaVersion,
					},
				},
				Mobile:   c.ChallengeFingerprint.Profile.Misc.Mobile,
				Platform: c.ChallengeFingerprint.Profile.Misc.Os,
			},
		},
		DR:   "",
		Inv:  c.Config.Invisible,
		Exec: false, //true, //false <--n
		Wn:   [][]int64{
			/*{
				int64(c.ChallengeFingerprint.Profile.Screen.AvailWidth),  // mt.Browser.width()   // ---> return window.innerWidth && window.document.documentElement.clientWidth ? Math.min(window.innerWidth, document.documentElement.clientWidth) : window.innerWidth || window.document.documentElement.clientWidth || document.body.clientWidth;
				int64(c.ChallengeFingerprint.Profile.Screen.AvailHeight), // mt.Browser.height()  // ---> return window.innerHeight || window.document.documentElement.clientHeight || document.body.clientHeight;
				1,                                       // mt.System.dpr()
				addTime(st, WnTime),                     // Date.now()
			},*/
		},
		WnMp: 0,
		Xy:   [][]int64{
			/*{
				0, // mt.Browser.scrollX(),  // ---> return window.pageXOffset !== undefined ? window.pageXOffset : WnTime.isCSS1 ? document.documentElement.scrollLeft : document.body.scrollLeft;
				0, // mt.Browser.scrollY(),  // ---> return window.pageYOffset !== undefined ? window.pageYOffset : WnTime.isCSS1 ? document.documentElement.scrollTop : document.body.scrollTop;
				int64(c.ChallengeFingerprint.Profile.Screen.AvailWidth) / (int64(c.ChallengeFingerprint.Profile.Screen.AvailWidth) * 2), // document.documentElement.clientWidth / mt.Browser.width(),
				addTime(st, WnTime), // Date.now()
			},*/
		},
		XyMp: 0,
		Mm:   MmPath,
		MmMp: calculateMeanPeriod(MmPath),
	}

	output := []byte{}

	switch m.IsCheck {
	case true:
		output, _ = json.Marshal(&CheckData{
			St:       st,
			Dct:      st,
			Mm:       CaptchaPath,
			MmMp:     calculateMeanPeriod(CaptchaPath),
			Md:       MdPath,
			MdMp:     calculateMeanPeriod(MdPath),
			Mu:       MuPath,
			MuMp:     calculateMeanPeriod(MuPath),
			TopLevel: topLevel,
			V:        1,
		})
	case false:
		widget := "0" + utils.RandomString(11)

		output, _ = json.Marshal(&GetData{
			St:       st,
			Mm:       CaptchaPath,
			MmMp:     calculateMeanPeriod(CaptchaPath),
			Md:       MdPath,
			MdMp:     calculateMeanPeriod(MdPath),
			Mu:       MuPath,
			MuMp:     calculateMeanPeriod(MuPath),
			TopLevel: topLevel,
			V:        1,

			Session: []string{},
			WidgetList: []string{
				widget,
			},
			WidgetID: widget,
			Href:     fmt.Sprintf("https://%s", c.Config.Domain),
			Prev: Prev{
				Escaped:          false,
				Passed:           false,
				ExpiredChallenge: false,
				ExpiredResponse:  false,
			},
		})
	}

	return string(output)
}
