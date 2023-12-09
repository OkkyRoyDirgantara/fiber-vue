<template>
  <div class="dashboard" style="margin-top: 80px">
    <div class="container">
      <div class="row">
        <div class="col-md-4">
          <div class="card">
            <div class="card-body">
              MAIN MENU
              <hr />
              <ul class="list-group">
                <router-link
                  :to="{ name: 'dashboard-home' }"
                  class="list-group-item text-dark text-decoration-none"
                  >Dashboard</router-link
                >
                <router-link
                  :to="{ name: 'products' }"
                  class="list-group-item text-dark text-decoration-none"
                  >Products</router-link
                >
                <li
                  @click="logout"
                  class="list-group-item text-dark text-decoration-none"
                  style="cursor: pointer"
                >
                  Logout
                </li>
              </ul>
            </div>
          </div>
        </div>
        <div class="col-md-8">
          <router-view> </router-view>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'DashboardView',
  data() {
    return {
      // state loggedIn with localStorage
      loggedIn: localStorage.getItem('loggedIn'),
      // state token
      token: localStorage.getItem('token'),
      // state user logged In
      user: {}
    }
  },
  created() {},
  methods: {
    logout() {
      try {
        localStorage.removeItem('loggedIn')

        // redirect
        this.$router.push({ name: 'login' })
      } catch (error) {
        console.error('Error redirecting to login:', error)
      }
    }
  },
  // check if the user is logged in or not
  mounted() {
    if (!this.loggedIn) {
      this.$router.push({ name: 'login' })
    }
  }
}
</script>
