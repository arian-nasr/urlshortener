<template>
  <div id="login" class="container">
    <form>
      <div class="form-group">
        <label for="email">Email address</label>
        <input type="email" class="form-control" id="email" placeholder="Enter email" v-model="email">
      </div>
      <div class="form-group">
        <label for="password">Password</label>
        <input type="password" class="form-control" id="password" placeholder="Password" v-model="password">
      </div>
      <button type="button" class="btn btn-primary" @click="submitform">Submit</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      email: '',
      password: '',
      token: ''
    }
  },
  created () {
    document.title = 'Login'
  },
  methods: {
    submitform: function () {
      const path = 'https://onebounce.me/api/auth/login'
      axios.post(path, {}, {
        auth: {
          username: this.email,
          password: this.password
        }
      })
        .then((res) => {
          this.token = res.data
          console.log(this.token)
        })
    }
  }
}
</script>
