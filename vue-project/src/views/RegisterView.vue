<template>
  <div class="container mt-5">
    <h2>Register</h2>
    <form @submit.prevent="registerUser">
      <div class="mb-3">
        <label for="name" class="form-label">Name:</label>
        <input type="text" class="form-control" v-model="userData.name" required />
      </div>

      <div class="mb-3">
        <label for="email" class="form-label">Email:</label>
        <input type="email" class="form-control" v-model="userData.email" required />
      </div>

      <div class="mb-3">
        <label for="password" class="form-label">Password:</label>
        <input type="password" class="form-control" v-model="userData.password" required />
      </div>

      <div class="mb-3">
        <label for="confirmPassword" class="form-label">Confirm Password:</label>
        <input type="password" class="form-control" v-model="userData.confirmPassword" required />
      </div>

      <button type="submit" class="btn btn-primary">Register</button>
    </form>

    <div v-if="errorMessage" class="alert alert-danger mt-3" role="alert">
      {{ errorMessage }}
    </div>
  </div>
</template>

<script>
import axios from 'axios'

let urlPath = import.meta.env.VITE_URL_BACKEND;

export default {
  name: 'RegisterView',
  data() {
    return {
      userData: {
        name: '',
        email: '',
        password: '',
        confirmPassword: ''
      },
      errorMessage: null
    }
  },

  methods: {
    registerUser() {
      // Validasi kata sandi
      if (this.userData.password !== this.userData.confirmPassword) {
        console.error('Password and Confirm Password do not match.')
        return
      }

      // Kirim data pendaftaran ke API
      axios
        .post(`${urlPath}/api/v1/user/register`, this.userData)
        .then((response) => {
          console.log('Registration successful:', response)
          // Reset formulir setelah pendaftaran berhasil
          this.userData = {
            name: '',
            email: '',
            password: '',
            confirmPassword: ''
          }

          this.$router.push({ name: 'login' })
        })
        .catch((error) => {
          console.error('Error registering user:', error)
          this.errorMessage = `Registration failed: ${error.response.data.data}`
        })
    }
  }
}
</script>
