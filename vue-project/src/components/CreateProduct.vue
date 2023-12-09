<template>
  <div class="mt-3">
    <div class="card mt-3">
      <div class="card-body">
        <form @submit.prevent="createProduct">
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

          <button type="submit" class="btn btn-primary">Create</button>
        </form>

        <div v-if="errorMessageCreate" class="alert alert-danger mt-3" role="alert">
          {{ errorMessageCreate }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

let urlPath = import.meta.env.VITE_URL_BACKEND;

export default {
  data() {
    return {
      createData: {
        name: '',
        code: '',
        stock: 1,
        description: '',
        status: true
      },
      errorMessageCreate: null
    }
  },

  methods: {
    createProduct() {
      axios
        .post(`${urlPath}/api/v1/product`, this.createData, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`
          }
        })
        .then(() => {
          this.$emit('productCreated', this.createData)
          this.createData.name = ''
          this.createData.code = ''
          this.createData.stock = 1
          this.createData.description = ''
          this.createData.status = true
          this.errorMessageCreate = null
        })
        .catch((err) => {
          console.log(err)
          this.errorMessageCreate = err.response.data.data
        })
    }
  }
}
</script>
