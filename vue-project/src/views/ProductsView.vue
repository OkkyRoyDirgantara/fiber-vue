<template>
  <div class="my-1 px-5 py-3">
    <!-- Add Product and Update Quantity buttons -->
    <div class="d-flex justify-content-between">
      <div>
        <router-link :to="{ name: 'products-create' }" class="btn btn-primary"
          >Create Product</router-link
        >
      </div>
      <div>
        <router-link :to="{ name: 'products-update' }" class="btn btn-warning"
          >Update Product Quantity</router-link
        >
      </div>
    </div>

    <!-- Router view for handling events -->
    <router-view
      @productCreated="handleProductCreated"
      @updateQuantity="handleUpdateQuantity"
    ></router-view>

    <!-- Product List Section -->
    <div class="mt-3 card px-5 py-3">
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
    async handleProductCreated() {
      await this.fetchProducts()
    },
    async handleUpdateQuantity() {
      await this.fetchProducts()
    }
  }
}
</script>
