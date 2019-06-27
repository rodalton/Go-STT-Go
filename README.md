# Go-STT-Go
In this repository we outline the steps involved in using the Watson SDK for Speech to Text using Go. We also include a specific example that outlines the approach to processing multiple audio files with STT using Go. 

## Environment Setup
To setup your environment, you first need to install a Go distribution. The appropriate distribution for your platform can be found here https://golang.org/doc/install

Once the installation of Go is complete, you can verify the installation using the instructions here https://golang.org/doc/install#testing 

Next, go ahead and install the Watson SDK for Go following the instructions in the next step. 

## Install the Watson SDK for Go
Instructions on how to install the Go SDK package for Watson can be found here https://github.com/watson-developer-cloud/go-sdk

Once you've installed the Watson SDK for Go, let's take a look at the examples available with this repository. 

## Using the Go-STT-Go examples
To use the examples made available in this repository, follow the steps below. 

1. Clone the source code from this repository using the folliwing command
`git clone https://github.com/rodalton/Go-STT-Go.git`

2. Edit the Go file you plan to run to include the API Key and service URL for you STT instance using a text editor. You'll find these values from the Service credentials tab of your Speech to Text service instance on IBM Cloud. Save the changes your Go file. 

3. From a terminal window, navigate to the directory that contains the Go file edited earlier. Issue the `go build` command to build an executable file that we'll run in the next step. 

4. Run the executable built in step 3 from the terminal window, for example issue the command `./RecognizeFile`. The executable will call the Watson Speech to Text API and print a transcript of the audio file to the terminal window. 




