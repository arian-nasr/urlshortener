<template>
    <form>
      <div class="form-group">
        <label for="exampleInputEmail1">Email address</label>
        <input type="email" class="form-control" id="exampleInputEmail1" placeholder="Enter email">
      </div>
      <div class="form-group">
        <label for="exampleInputPassword1">Password</label>
        <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
      </div>
      <button type="submit" class="btn btn-primary" v-onclick="submitform">Submit</button>
    </form>
</template>

<script>
export default {
  created () {
    document.title = 'Login'
  }
  methods: {
    submitform: function (event) {
      alert('Test!')
      console.log('Test!')
    }
  }
}
</script>
