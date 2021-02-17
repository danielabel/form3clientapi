package form3clientapi

import "testing"

func TestSettingURL(t *testing.T) {

	t.Run("get base URL from defaults", func(t *testing.T) {
		resetDomain()
		expected := "http://localhost:8080/v1"
		url := getBaseUrl()
		if url != expected {
			t.Errorf("expected baseURL from defaults: %s, got %s", expected, url)
		}
	})

	t.Run("get base URL with custom domain", func(t *testing.T) {
		resetDomain()
		expected := "http://customhost:8080/v1"
		setDomain("customhost", "")
		url := getBaseUrl()
		if url != expected {
			t.Errorf("expected baseURL from defaults: %s, got %s", expected, url)
		}
	})

	t.Run("get base URL with custom port", func(t *testing.T) {
		resetDomain()
		expected := "http://localhost:9090/v1"
		setDomain("", "9090")
		url := getBaseUrl()
		if url != expected {
			t.Errorf("expected baseURL from defaults: %s, got %s", expected, url)
		}
	})
}