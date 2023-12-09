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

        axios
          .post('http://localhost:3000/api/v1/auth/login', {
            email: this.user.email,
            password: this.user.password
          })
          .then((res) => {
            //debug user login
            console.log(res.data)

            if (res.data.status == 'success') {
              //set localStorage
              localStorage.setItem('loggedIn', 'true')

              //set localStorage Token
              localStorage.setItem('token', res.data.token)

              //change state
              this.loggedIn = true

              return this.$router.push({ name: 'dashboard-home' })
            } else {
              //set state login failed
              this.loginFailed = true
            }
          })
          .catch((error) => {
            this.loginFailed = true
            console.log(error)
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
