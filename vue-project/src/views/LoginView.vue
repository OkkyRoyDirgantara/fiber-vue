<template>
  <div class="login">
    <div class="container mt-5">
      <div class="row justify-content-center">
        <div class="col-md-5">
          <div v-if="loginFailed" class="alert alert-danger">Email atau Password Anda salah.</div>
          <div class="card">
            <div class="card-body">
              <h2 class="mb-4">LOGIN</h2>
              <hr />

              <form @submit.prevent="login">
                <div class="mb-3">
                  <label for="email" class="form-label">EMAIL</label>
                  <input
                    type="email"
                    class="form-control"
                    v-model="user.email"
                    placeholder="Masukkan Email"
                    required
                  />
                  <div v-if="validation.email" class="mt-2 alert alert-danger">Masukkan Email</div>
                </div>

                <div class="mb-3">
                  <label for="password" class="form-label">PASSWORD</label>
                  <input
                    type="password"
                    class="form-control"
                    v-model="user.password"
                    placeholder="Masukkan Password"
                    required
                  />
                  <div v-if="validation.password" class="mt-2 alert alert-danger">
                    Masukkan Password
                  </div>
                </div>

                <button type="submit" class="btn btn-primary">LOGIN</button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import BackendService from '../router/backend.ts'

let urlPath = import.meta.env.VITE_URL_BACKEND

export default {
  name: 'LoginView',

  data() {
    return {
      //state loggedIn with localStorage
      loggedIn: localStorage.getItem('loggedIn'),
      //state token
      token: localStorage.getItem('token'),
      //state user
      user: [],
      //state validation
      validation: [],
      //state login failed
      loginFailed: false
    }
  },
  methods: {
    login() {
      if (this.user.email && this.user.password) {
        //debug cookie
        // console.log(response)

        let api = new BackendService()
        api
          .login({ email: this.user.email, password: this.user.password })
          .then((response) => {
            //debug cookie
            //set localStorage
            localStorage.setItem('loggedIn', 'true')

            //set localStorage Token
            localStorage.setItem('token', response.token)

            //change state
            this.loggedIn = true

            //redirect to dashboard
            this.$router.push({ name: 'dashboard-home' })
          })
          .catch((error) => {
            console.error('Error login:', error)
            this.loginFailed = true
          })
      }

      this.validation = []

      if (!this.user.email) {
        this.validation.email = true
      }

      if (!this.user.password) {
        this.validation.password = true
      }
    }
  },

  //check user already logged in
  mounted() {
    if (this.loggedIn) {
      return this.$router.push({ name: 'dashboard-home' })
    }
  }
}
</script>
