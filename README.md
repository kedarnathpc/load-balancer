
<div align="center">

  <h1 align="center">Load Balancer using Golang</h1>

</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This project implements a simple Load Balancer using Golang, reverse proxy, and Docker. The load balancer distributes incoming network traffic across multiple backend servers to ensure optimal resource utilization and maintain high availability.


### Built With

* [![Golang][Golang.shield]][Golang.url]
* [![Docker][Docker.shield]][Docker.url]


<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

Following are the prerequsites for the project.
* Golang
* Docker
* Docker-compose

### Installation

_Below are the steps of installation._

1. Clone the repo
   ```sh
   git clone https://github.com/kedarnathpc/load-balancer
   ```
2. Build the Docker images
    ```
    docker-compose build    
    ```
3. Run the Docker containers 
   ```sh
   docker-compose up -d
   ```


<!-- USAGE EXAMPLES -->
## Usage

* Access the load balancer at http://localhost:8080.
* Add or modify backend servers in the files as needed.



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[Golang.shield]: https://img.shields.io/badge/Golang-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Golang.url]: https://golang.org/
[Docker.shield]: https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white
[Docker.url]: https://www.docker.com/

