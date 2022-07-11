
  <h1>One Piece Api</h1>
   <sub> Basic API developed in Go & GraphQL</sub>
<br />
<br />
<div align="center">
  <a>
  
  ![One_Piece_Anime_Logo](https://user-images.githubusercontent.com/38867931/178217876-43b54feb-15f1-4709-9f81-15d836e62f18.png)

  </a>
</div>



## About The Project
In my opinion the best way to study new techologies is to make projects, so while studying Go and GraphQL i decided to create a basic One Piece (famous manga) API that returns the Strawhats crew information, and give you the possibility to create new crews. 


### Built With
* [Go](https://go.dev/)
* [GraphQL](https://graphql.org/)

<!-- GETTING STARTED -->
## Getting Started
Start the server : 
  ```sh
  go run .
  ```

## Usage

Using GraphQL playground in http://localhost:8080/ (in my case) you can try this Queries and Mutations

* Get all the strawhats crew 
```graphql
  query findPirates{
  strawhats{
    name
    image
    bounty
    crew
    id
  }
}
  ```
  
  
<details>
  <summary>Output :</summary>
  
  ```graphql
 {
  "data": {
    "strawhats": [
      {
        "name": "Monkey D. Luffy",
        "image": "https://rb.gy/xz0x9h",
        "bounty": "1,500,000,000",
        "crew": "Strawhats",
        "id": "01"
      },
      {
        "name": "Roronoa Zoro",
        "image": "https://rb.gy/imkdhx",
        "bounty": "320,000,000",
        "crew": "Strawhats",
        "id": "02"
      },
      {
        "name": "Nami",
        "image": "https://rb.gy/dncyia",
        "bounty": "66,000,000",
        "crew": "Strawhats",
        "id": "03"
      },
      {
        "name": "God Usop",
        "image": "https://rb.gy/xa68o1",
        "bounty": "200,000,000",
        "crew": "Strawhats",
        "id": "04"
      },
      {
        "name": "Vinsmoke Sanji ",
        "image": "https://rb.gy/sofr7o",
        "bounty": "330,000,000",
        "crew": "Strawhats",
        "id": "05"
      },
      {
        "name": "Tony Tony Chopper",
        "image": "https://rb.gy/oh2nmv",
        "bounty": "100",
        "crew": "Strawhats",
        "id": "06"
      },
      {
        "name": "Nico Robin",
        "image": "https://rb.gy/5puvys",
        "bounty": "130,000,000",
        "crew": "Strawhats",
        "id": "07"
      },
      {
        "name": "Brook",
        "image": "https://rb.gy/5kelik",
        "bounty": "83,000,000",
        "crew": "Strawhats",
        "id": "08"
      },
      {
        "name": "Jimbei",
        "image": "https://rb.gy/mafbmm",
        "bounty": "83,000,000",
        "crew": "Strawhats",
        "id": "09"
      },
      {
        "name": "Neferutari Bibi",
        "image": "https://rb.gy/zddixu",
        "bounty": "0",
        "crew": "Strawhats",
        "id": "11"
      }
    ]
  }
}
  ```
</details>

---


* Create your crew, adding new pirates 
```graphql
 mutation createPirate {
  createPirate(input: { name: "Marco", pirateId: "1",bounty:" 1.374.000.000",crew:"Whitebeard crew",image:"null"}) {
	  name
	  id
  }
}
  ```
<details>
  <summary>Output :</summary>
  
  ```graphql
 {
  "data": {
    "createPirate": {
      "name": "Marco",
      "id": "T8674665223082153551",
    }
  }
}
  ```
</details>

---

* Return the created pirates 
```graphql
 query findPirates{
  pirates{
    name
    id
    bounty
    crew
    image //optional
  }
}
  ```
<details>
  <summary>Output :</summary>
  
  ```graphql
 {
  "data": {
    "pirates": [
      {
        "name": "Edward Newgate",
        "id": "T5577006791947779410",
        "bounty": "5,046,000,000",
        "crew": "Whitebeard crew",
        "image": null
      },
      {
        "name": "Marco",
        "id": "T8674665223082153551",
        "bounty": "1.374.000.000",
        "crew": "Whitebeard crew",
        "image": "null" 
      }
    ]
  }
}
  ```
</details>

---

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**. 

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again! ❤️

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>


## Contact

Matteo Leonesi - [Linkedin](https://www.linkedin.com/in/matteo-leonesi-228867138/) - matteo.leonesi@gmail.com

Project Link: [GraphQL-Go-onepieceAPI](https://github.com/MatteoLeonesi/GraphQL-Go-onepieceAPI)

<p align="right">(<a href="#top">back to top</a>)</p>


