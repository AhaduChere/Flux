import express from "express";
const app = express();
app.use(express.json());

app.get("/api/scores", (req, res) => {
  res.json({
    scores: [
      { user: "Jamie", score: 984 },
      { user: "Reza", score: 761 },
    ],
    total: 2,
  });
});

app.post("/api/users", (req, res) => {
  const { name, role } = req.body;
  res.status(201).json({ id: 3, name, role });
});

app.listen(3000, () => console.log("Server running on 3000"));
