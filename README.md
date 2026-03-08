<p align="center">
  <img src="assets/Logo.svg" alt="Flux Logo" width="120">
</p>
<h1 align="center">Flux</h1>
<p align="center"><em>A Command-Line API Client I made similar to Postman</em></p>
<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=for-the-badge" alt="Go">
  <img src="https://img.shields.io/badge/CLI-2D3748?style=for-the-badge&logo=gnubash&logoColor=white">
  <img src="https://img.shields.io/badge/API-FF6B6B?style=for-the-badge&logo=fastapi&logoColor=white">
</p>

<h4>How to Use</h4>
<ul>
  <li>Define API endpoints in a Fluxfile at the root of the project</li>
  <li>Run any endpoint by name via the command line</li>
</ul>

<h4>Example Fluxfile Setup</h4> 

```toml
url = "http://localhost:3000"

[req.scores]
path = "/api/scores"
method = "GET"

[req.createUser]
path = "/api/users"
method = "POST"
body = { name = "Jamie", role = "admin" }
```

<h4>Example Use</h4> 

```bash
$ flux scores
✓ 200 OK · 42ms
{
  "scores": [
    { "user": "Jamie", "score": 984 },
    { "user": "Reza",  "score": 761 }
  ],
  "total": 2
}

$ flux createUser
✓ 201 Created · 103ms
{
  "id": 3,
  "name": "Jamie",
  "role": "admin"
}
```

<h4>Future Plans</h4>
<ul>
  <li>Add support for custom request headers </li>
  <li>Add support for query parameters in the Fluxfile</li>
</ul>
