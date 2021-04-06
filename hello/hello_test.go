package hello

import "testing"

func TestHello(t *testing.T) {
    verifyMessageCorrect := func(t *testing.T, result, expected string) {
        t.Helper()
        if result != expected {
            t.Errorf("result '%s', expected '%s'", result, expected)
        }
    }

    t.Run("Say hello to new people", func(t *testing.T) {
        result := Hello("Edu", "")
        expected := "Hello, Edu"
    
        verifyMessageCorrect(t, result, expected)
    })

    t.Run("Say 'Hello, World' when an empty string is passed", func(t *testing.T) {
        result := Hello("", "")
        expected := "Hello, World"

        verifyMessageCorrect(t, result, expected)
    })

    t.Run("in Spanish", func(t *testing.T) {
        result := Hello("Edu", "spanish")
        expected := "Hola, Edu"

        verifyMessageCorrect(t, result, expected)
    })

    t.Run("in French", func(t *testing.T) {
        result := Hello("Edu", "french")
        expected := "Bonjour, Edu"

        verifyMessageCorrect(t, result, expected)
    })
}