

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">OPN Generator</h3>
  <p align="center">
    OPN Generator is a tool to build a new application by providing API specification and Database schema.  
    <br />
    <br />
    <a href="https://github.com/opn-ooo/opn-generator/issues/new">Report Bug</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#project-structure">Project Structure</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#requirements">Requirements</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project


### Project Structure

```
📦opn-generator
 ┣ 📂cmd
 ┣ 📂config
 ┣ 📂internal
 ┃ ┣ 📂generators
 ┃ ┣ 📂handlers
 ┃ ┣ 📂services
 ┣ 📂out
 ┃ ┣ 📂db
 ┃ ┗ 📂temp
 ┣ 📂pkg
 ┃ ┣ 📂database_schema
 ┃ ┣ 📂files
 ┃ ┣ 📂open_api_spec
 ┃ ┣ 📂template
 ```
This tool is developed by following [Standard Go project layout](https://github.com/golang-standards/project-layout) 

<!-- GETTING STARTED -->
## Getting Started

You can use MacOS, Linux or Windows as operating system to use this tool. To get a local copy up and running follow these simple steps.

### Prerequisites
1. Make sure [Git](https://git-scm.com/downloads) is installed. To verify that you've installed Git by opening a command prompt and typing the following command that  prints the installed version of Git.
    ```sh
    $ git version
    ```
2. Make sure [GO](https://golang.org/doc/install) is installed. To verify that you've installed Go by opening a command prompt and typing the following command that  prints the installed version of Go.
    ```sh
    $ go version
    ```
3. Prepare API spec ([OpenAPISpec](https://spec.openapis.org/oas/latest.html) ver.3 YAML file).
4. Prepare DB Schema design ([PlantUML](https://plantuml.com) file).

### Installation

1. Clone - Go to a folder where you would like to clone this tool and run following.
   ```sh
   $ git clone https://github.com/opn-ooo/opn-generator.git
   ```
2. Build - Go to the opn-generator directory and run following. 
   ```sh
   $ go build
   ```

<!-- USAGE EXAMPLES -->
## Usage

To use OPN Generator after the installation succeed,  please use the following command that will build a new application
   ```sh
   $ ./opn-generator new [projectName] --database [/path/to/plantUMLFile] --api [/path/to/apiSpecYAMLFile]
   ```

   Here, \
   [projectName] is name of the project \
   --database is the file path of database plantuml file, \
   --api is the file path of API specification file. 
    

_To check sample files of API specification and Database schema, have a look over [here](sample/)_

<!-- ROADMAP -->
## Roadmap

The roadmap is useful for planning large pieces of work several months in advance at the Epic level within a single project. Simple planning and dependency management features help teams visualize and manage work better together. 


See the [open issues](https://github.com/opn-ooo/opn-generator/issues?q=is%3Aopen+is%3Aissue) for a list of proposed features (and known issues).



<!-- REQUIREMENTS -->
## Requirements

* Go^1.15

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [Dig](https://pkg.go.dev/go.uber.org/dig)
* [Cobra](https://cobra.dev)
* [Testify](https://github.com/stretchr/testify) \
And many more


