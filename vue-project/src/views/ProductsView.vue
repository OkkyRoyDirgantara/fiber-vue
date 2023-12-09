<template>
  <div class="container mt-5">
    <h2 class="mt-5">Create Product</h2>
    <form @submit.prevent="createProduct" class="mt-3">
      <div class="mb-3">
        <label for="name" class="form-label">Product Name:</label>
        <input v-model="createData.name" type="text" class="form-control" required />
      </div>

      <div class="mb-3">
        <label for="code" class="form-label">Product Code:</label>
        <input v-model="createData.code" type="text" class="form-control" required />
      </div>

      <div class="mb-3">
        <label for="stock" class="form-label">Product Stock:</label>
        <input v-model="createData.stock" type="number" class="form-control" min="1" required />
      </div>

      <div class="mb-3">
        <label for="description" class="form-label">Product Description:</label>
        <textarea
          v-model="createData.description"
          class="form-control"
          rows="3"
          required
        ></textarea>
      </div>

      <div class="mb-3">
        <label for="status" class="form-label">Product Status:</label>
        <select v-model="createData.status" class="form-select" required>
          <option value="true">Active</option>
          <option value="false">Inactive</option>
        </select>
      </div>

      <button type="submit" class="btn btn-primary">Create Product</button>
    </form>
    <div v-if="errorMessageCreate" class="alert alert-danger mt-3" role="alert">
      {{ errorMessageCreate }}
    </div>

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

    <div>
      <h2>Product List</h2>
      <div class="row">
        <div class="col-md-4" v-for="product in products" :key="product.Code">
          <div class="card my-2">
            <div class="card-body">
              <h3 class="card-title">{{ product.Name }}</h3>
              <ul class="list-group list-group-flush">
                <li class="list-group-item"><strong>Code:</strong> {{ product.Code }}</li>
                <li class="list-group-item"><strong>Stock:</strong> {{ product.Stock }}</li>
                <li class="list-group-item">
                  <strong>Description:</strong> {{ product.Description }}
                </li>
                <li class="list-group-item"><strong>Status:</strong> {{ product.Status }}</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      products: [],
      updateData: {
        code: '',
        action: 'add',
        amount: 1
      },
      createData: {
        name: '',
        code: '',
        stock: 1,
        description: '',
        status: true
      },
      errorMessage: null,
      errorMessageCreate: null
    }
  },
  mounted() {
    this.fetchProducts()
  },
  methods: {
    async fetchProducts() {
      try {
        const response = await axios.get('http://127.0.0.1:3000/api/v1/product', {
          headers: {
            Authorization: 'Bearer ' + localStorage.getItem('token')
          }
        })
        this.products = response.data.data
      } catch (error) {
        console.error('Error fetching products:', error)
      }
    },
    async updateQuantity() {
      try {
        await axios.patch('http://127.0.0.1:3000/api/v1/product', this.updateData, {
          headers: {
            Authorization: 'Bearer ' + localStorage.getItem('token')
          }
        })
        this.fetchProducts()
        this.updateData = {
          code: '',
          action: 'add',
          amount: 1
        }
        this.errorMessage = null
      } catch (error) {
        console.error('Error updating quantity:', error)
        this.errorMessage = 'Error updating quantity: ' + error.response.data.message
      }
    },
    async createProduct() {
      try {
        await axios.post('http://127.0.0.1:3000/api/v1/product', this.createData, {
          headers: {
            Authorization: 'Bearer ' + localStorage.getItem('token')
          }
        })
        this.fetchProducts()
        this.createData = {
          name: '',
          code: '',
          stock: 1,
          description: '',
          status: true
        }
        this.errorMessageCreate = null
      } catch (error) {
        console.error('Error creating product:', error)
        this.errorMessageCreate = 'Error creating product: ' + error.response.data.message
      }
    }
  }
}
</script>
