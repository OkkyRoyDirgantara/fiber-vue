<template>
  <div class="mt-3">
    <div class="card mt-3">
      <div class="card-body">
        <form @submit.prevent="updateQuantity">
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
            <input
              v-model="updateData.amount"
              type="number"
              class="form-control"
              min="1"
              required
            />
          </div>

          <button type="submit" class="btn btn-primary">Update</button>
        </form>

        <div v-if="errorMessage" class="alert alert-danger mt-3" role="alert">
          {{ errorMessage }}
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
      errorMessage: null
    }
  },
  methods: {
    async updateQuantity() {
      axios
        .patch('http://localhost:3000/api/v1/product', this.updateData, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`
          }
        })
        .then(() => {
          this.$emit('updateQuantity', this.updateData)
          this.updateData.code = ''
          this.updateData.action = 'add'
          this.updateData.amount = 1
          this.errorMessage = null
        })
        .catch((err) => {
          console.log(err)
          this.errorMessage = err.response.data.message
        })
    }
  }
}
</script>
