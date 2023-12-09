<template>
  <div class="container mt-5">
    <h2 class="mt-5">Update Product Quantity</h2>
    <form @submit.prevent="updateQuantity" class="mt-3">
      <div class="mb-3">
        <label for="code" class="form-label">Product Code:</label>
        <input v-model="updateData.code" type="text" class="form-control" required />
      </div>

      <div class="mb-3">
        <label for="action" class="form-label">Action:</label>
        <select v-model="updateData.action" class="form-select" required>
          <option value="add">Add</option>
          <option value="subtract">Subtract</option>
        </select>
      </div>

      <div class="mb-3">
        <label for="amount" class="form-label">Amount:</label>
        <input v-model="updateData.amount" type="number" class="form-control" min="1" required />
      </div>

      <button type="submit" class="btn btn-primary">Update Quantity</button>
    </form>
    <div v-if="errorMessage" class="alert alert-danger mt-3" role="alert">
    {{ errorMessage }}
  </div>

    <h2>Product List</h2>
    <div class="card my-2" v-for="product in products" :key="product.Code">
    <ul class="list-group list-group-flush">
      <li class="list-group-item">
        <div class="card-body">
          <h5 class="card-title">{{ product.Name }}</h5>
          <p class="card-text">Code: {{ product.Code }}</p>
          <p class="card-text">Stock: {{ product.Stock }}</p>
        </div>
      </li>
    </ul>
  </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      products: [],
      updateData: {
        code: '',
        action: 'add',
        amount: 1,
      },
      errorMessage: null,
    };
  },
  mounted() {
    this.fetchProducts();
  },
  methods: {
    async fetchProducts() {
      try {
        const response = await axios.get('http://127.0.0.1:3000/api/v1/product', {
          headers: {
            Authorization: 'Bearer ' + localStorage.getItem('token'),
          },
        });
        this.products = response.data.data;
      } catch (error) {
        console.error('Error fetching products:', error);
      }
    },
    async updateQuantity() {
      try {
        await axios.patch('http://127.0.0.1:3000/api/v1/product', this.updateData,{
          headers: {
            Authorization: 'Bearer ' + localStorage.getItem('token'),
          },
        });
        this.fetchProducts();
        this.errorMessage = null;
      } catch (error) {
        console.error('Error updating quantity:', error);
        this.errorMessage = 'Error updating quantity: ' + error.response.data.message;
      }
    },
  },
};
</script>
