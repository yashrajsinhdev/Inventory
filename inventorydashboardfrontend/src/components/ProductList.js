"use client";
import { useState } from "react";
import {
  Paper,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  IconButton,
  TextField,
} from "@mui/material";
import {
  EditOutlined,
  DeleteOutlined,
  SaveOutlined,
  CloseOutlined,
} from "@ant-design/icons";

export default function ProductList({ products, onChange }) {
  const [editId, setEditId] = useState(null);
  const [editName, setEditName] = useState("");
  const [editQuantity, setEditQuantity] = useState("");

  const startEdit = (product) => {
    setEditId(product.id);
    setEditName(product.name);
    setEditQuantity(product.quantity);
  };

  const cancelEdit = () => {
    setEditId(null);
    setEditName("");
    setEditQuantity("");
  };

  const saveEdit = async (id) => {
    await fetch(`http://localhost:8080/product/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        name: editName,
        quantity: parseInt(editQuantity),
      }),
    });
    cancelEdit();
    onChange();
  };

  const deleteProduct = async (id) => {
    await fetch(`http://localhost:8080/product/${id}`, { method: "DELETE" });
    onChange();
  };

  return (
    <Paper sx={{ p: 2 }}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell width="50%">Name</TableCell>
            <TableCell width="30%">Quantity</TableCell>
            <TableCell width="20%">Actions</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {products.length === 0 ? (
            <TableRow>
              <TableCell colSpan={3} align="center">
                No products found.
              </TableCell>
            </TableRow>
          ) : (
            products.map((product) =>
              editId === product.id ? (
                <TableRow key={product.id}>
                  <TableCell>
                    <TextField
                      value={editName}
                      onChange={(e) => setEditName(e.target.value)}
                      size="small"
                      fullWidth
                    />
                  </TableCell>
                  <TableCell>
                    <TextField
                      type="number"
                      value={editQuantity}
                      onChange={(e) => setEditQuantity(e.target.value)}
                      size="small"
                      fullWidth
                    />
                  </TableCell>
                  <TableCell>
                    <IconButton
                      color="primary"
                      onClick={() => saveEdit(product.id)}
                    >
                      <SaveOutlined />
                    </IconButton>
                    <IconButton color="error" onClick={cancelEdit}>
                      <CloseOutlined />
                    </IconButton>
                  </TableCell>
                </TableRow>
              ) : (
                <TableRow key={product.id}>
                  <TableCell>{product.name}</TableCell>
                  <TableCell>{product.quantity}</TableCell>
                  <TableCell>
                    <IconButton
                      color="primary"
                      onClick={() => startEdit(product)}
                    >
                      <EditOutlined />
                    </IconButton>
                    <IconButton
                      color="error"
                      onClick={() => deleteProduct(product.id)}
                    >
                      <DeleteOutlined />
                    </IconButton>
                  </TableCell>
                </TableRow>
              )
            )
          )}
        </TableBody>
      </Table>
    </Paper>
  );
}
