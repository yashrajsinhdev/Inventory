"use client";
import { useState, useEffect } from "react";
import {
  Box,
  Typography,
  Stack,
  TextField,
  Select,
  MenuItem,
  Paper,
} from "@mui/material";
import AddProduct from "../components/AddProduct";
import ProductList from "../components/ProductList";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export default function Home() {
  const [products, setProducts] = useState([]);
  const [filter, setFilter] = useState("");
  const [sort, setSort] = useState("");

  const fetchProducts = async () => {
    const params = new URLSearchParams();
    if (filter) params.append("name", filter);
    if (sort) params.append("sort", sort);
    const res = await fetch(`${API_URL}/products?${params}`);
    const data = await res.json();
    setProducts(data);
  };

  useEffect(() => {
    fetchProducts();
  }, [filter, sort]);

  return (
    <Box>
      <Typography variant="h4" gutterBottom textAlign="center">
        Inventory Dashboard
      </Typography>
      <AddProduct onAdd={fetchProducts} />
      <Paper sx={{ p: 2, mb: 2 }}>
        <Stack direction={{ xs: "column", sm: "row" }} spacing={2}>
          <TextField
            label="Filter by name"
            value={filter}
            onChange={(e) => setFilter(e.target.value)}
            fullWidth
          />
          <Select
            value={sort}
            onChange={(e) => setSort(e.target.value)}
            displayEmpty
            fullWidth
          >
            <MenuItem value="">Sort by</MenuItem>
            <MenuItem value="quantity">Quantity (Low to High)</MenuItem>
            <MenuItem value="name">Name (A-Z)</MenuItem>
          </Select>
        </Stack>
      </Paper>
      <ProductList products={products} onChange={fetchProducts} />
    </Box>
  );
}
