package main

import (
	"fmt"
	"io/ioutil"
	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
	"os"
)

func main() {

	files, err := ioutil.ReadDir("audio/")
	if err != nil {
			fmt.Println(err)
	}

  for _, f := range files {
		// Instantiate the Watson Speech To Text service
		service, serviceErr := speechtotextv1.
			NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
				URL:       "<URL>",
				IAMApiKey: "<API_KEY>",
			})

		// Check successful instantiation
		if serviceErr != nil {
			panic(serviceErr)
		}

		/* RECOGNIZE */
		fmt.Println(f.Name())
		pwd, _ := os.Getwd()
		// Open file with mp3 to recognize
		audio, audioErr := os.Open(pwd + "../audio/" + f.Name())
		if audioErr != nil {
			panic(audioErr)
		}

		// Create a new RecognizeOptions for ContentType "audio/mp3"
		recognizeOptions := service.
			NewRecognizeOptions(audio).
			SetContentType(speechtotextv1.RecognizeOptions_ContentType_AudioWav)

	  // Call the speechToText Recognize method
	  response, responseErr := service.Recognize(recognizeOptions)

		// Check successful call
		if responseErr != nil {
			panic(responseErr)
		}

		// Cast recognize.Result to the specific dataType returned by Recognize
		// NOTE: most methods have a corresponding Get<methodName>Result() function
		recognizeResult := service.GetRecognizeResult(response)
		//defer wg.Done()

		// Check successful casting
		if recognizeResult != nil {
			core.PrettyPrint(recognizeResult, "Recognize")
		}
  }
}
