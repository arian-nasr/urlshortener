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
export default {
  data () {
    return {
      email: '',
      password: ''
    }
  },
  created () {
    document.title = 'Login'
  },
  methods: {
    submitform: function () {
      alert(this.email)
    }
  }
}
</script>
