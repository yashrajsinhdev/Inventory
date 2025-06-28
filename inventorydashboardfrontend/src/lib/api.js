const API_URL = "http://localhost:8080";

export const fetchProducts = async (filter = "", sort = "") => {
  const params = new URLSearchParams();
  if (filter) params.append("name", filter);
  if (sort) params.append("sort", sort);

  const res = await fetch(`${API_URL}/products?${params}`);
  return await res.json();
};

export const addProduct = async (product) => {
  const res = await fetch(`${API_URL}/product`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(product),
  });
  return await res.json();
};

export const updateProduct = async (id, product) => {
  await fetch(`${API_URL}/product/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(product),
  });
};

export const deleteProduct = async (id) => {
  await fetch(`${API_URL}/product/${id}`, {
    method: "DELETE",
  });
};
