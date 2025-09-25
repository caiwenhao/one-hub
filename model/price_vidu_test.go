package model

import "testing"

func TestViduPriceConversion(t *testing.T) {
	prices := GetDefaultPrice()

	cases := []struct {
		model   string
		credits float64
	}{
		{"vidu-img2video-viduq2-pro-2s-720p", 3},
		{"vidu-img2video-viduq2-turbo-8s-1080p", 11},
		{"vidu-reference2video-vidu2.0-4s-720p", 8},
		{"vidu-text2video-vidu1.5-4s-1080p-anime", 20},
	}

	const rate = 0.3125

	for _, tc := range cases {
		t.Run(tc.model, func(t *testing.T) {
			var price *Price
			for _, p := range prices {
				if p.Model == tc.model {
					price = p
					break
				}
			}
			if price == nil {
				t.Fatalf("price for %s not found", tc.model)
			}
			expected := tc.credits * rate
			if price.Input != expected || price.Output != expected {
				t.Fatalf("price mismatch for %s: got %f/%f want %f", tc.model, price.Input, price.Output, expected)
			}
		})
	}
}
