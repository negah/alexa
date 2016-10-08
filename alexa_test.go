package alexa

import (
  "testing"
)

func TestGlobalRank(t *testing.T) {
  url := "google.com"
  rank := GlobalRank(url)

  if rank != "1" {
    t.Fatalf("Wrong global rank!")
  }
}

func TestCountryRank(t *testing.T) {
  url := "google.com"

  rank, country_name, country_code := CountryRank(url)

  if rank != "1" && country_name != "United States" && country_code != "US" {
    t.Fatalf("Wrong country rank!")
  }
}
