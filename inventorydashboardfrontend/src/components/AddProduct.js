"use client";
import { useState } from "react";
import { Paper, Box, Stack, TextField, Button } from "@mui/material";

export default function AddProduct({ onAdd }) {
  const [name, setName] = useState("");
  const [quantity, setQuantity] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!name.trim() || !quantity || quantity <= 0) {
      setError("Please enter a valid name and a positive quantity.");
      return;
    }
    setError("");
    await fetch("http://localhost:1111/product", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, quantity: parseInt(quantity) }),
    });
    setName("");
    setQuantity("");
    onAdd();
  };

  return (
    <Paper sx={{ p: 2, mb: 2 }}>
      <Box component="form" onSubmit={handleSubmit}>
        <Stack direction={{ xs: "column", sm: "row" }} spacing={2}>
          <TextField
            label="Product Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            error={!!error && !name.trim()}
            fullWidth
          />
          <TextField
            label="Quantity"
            type="number"
            value={quantity}
            onChange={(e) => setQuantity(e.target.value)}
            error={!!error && (!quantity || quantity <= 0)}
            fullWidth
          />
          <Button type="submit" variant="contained">
            Add
          </Button>
        </Stack>
        {error && (
          <Box sx={{ color: "error.main", mt: 1, textAlign: "center" }}>
            {error}
          </Box>
        )}
      </Box>
    </Paper>
  );
}
