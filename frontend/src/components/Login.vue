<template>
  <div id="login" class="container">
    <form>
      <div class="form-group">
        <label for="exampleInputEmail1">Email address</label>
        <input type="email" class="form-control" id="exampleInputEmail1" placeholder="Enter email">
      </div>
      <div class="form-group">
        <label for="exampleInputPassword1">Password</label>
        <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
      </div>
      <button type="button" class="btn btn-primary" @click="submitform">Submit</button>
    </form>
  </div>
</template>

<script>
export default {
  created () {
    document.title = 'Login'
  },
  methods: {
    submitform: function () {
      alert('Test!')
      console.log('Test!')
    }
  }
}
</script>
