package alexa

import (
  "testing"
)

func TestGlobalRank(t *testing.T) {
  url := "google.com"
  rank, err := GlobalRank(url)
  if err != nil {
    t.Fatalf("An error occured with getting rank")
  }
  if rank != "1" {
    t.Fatalf("Wrong global rank!")
  }
}

func TestCountryRank(t *testing.T) {
  url := "google.com"

  rank, country_name, country_code, err := CountryRank(url)

  if err != nil {
    t.Fatalf("An error occured with getting rank")
  }

  if rank != "1" && country_name != "United States" && country_code != "US" {
    t.Fatalf("Wrong country rank!")
  }
}
