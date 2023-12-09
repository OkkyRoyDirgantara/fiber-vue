<template>
  <div class="card">
    <div class="card-body">
      DASHBOARD
      <hr />
      Selamat Datang <strong>{{ user.name }}</strong>
    </div>
  </div>
</template>
<script>
import axios from 'axios'
export default {
  name: 'DashboardHomeView',
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
  created() {
    axios
      .get(`http://localhost:8000/api/user`, { headers: { Authorization: 'Bearer ' + this.token } })
      .then((response) => {
        this.user = response.data // assign response to state user
      })
      .catch((error) => {
        console.error('Error fetching user data:', error)
      })
  },
  mounted() {
    if (!this.loggedIn) {
      this.$router.push({ name: 'login' })
    }
  }
}
</script>
