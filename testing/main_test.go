package main

import "testing"

type TestCase struct {
	price       float64
	discount    float64
	expected    float64
	expectError bool
	desc        string
}

func TestDiscountPrice(t *testing.T) {
	testCases := []TestCase{
		{100, 0.0, 100, false, "discount  0%"},
		{100, -1.0, 0, true, "discount less than 0%"},
		{1000, 0.6, 400, false, "discount 60%"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result , err := DiscountedPrice(tc.price, tc.discount)
			if tc.expectError && err == nil {
				t.Errorf("DiscountedPrice(%.2f , %.2f) must return an error", tc.price, tc.discount)
			}
			if !tc.expectError && err != nil {
				t.Errorf("DiscountedPrice(%.2f , %.2f) must not return an error", tc.price, tc.discount)
			}
			if  !tc.expectError && result  != tc.expected {
				t.Errorf("DiscountedPrice(%.2f , %.2f) must return %.2f", tc.price, tc.discount, tc.expected)

			}

		})
	}

}
