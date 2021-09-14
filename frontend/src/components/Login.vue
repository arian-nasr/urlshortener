<template>
  <div class="modal modal-signin position-static d-block py-5" tabindex="-1" role="dialog" id="modalSignin">
    <div class="modal-dialog" role="document">
      <div class="modal-content rounded-6 shadow">
        <div class="modal-header p-5 pb-4 border-bottom-0">
          <!-- <h5 class="modal-title">Modal title</h5> -->
          <h2 class="fw-bold mb-0">Sign In</h2>
        </div>

        <div class="modal-body p-5 pt-0">
          <form class="">
            <div class="form-floating mb-3">
              <input type="email" class="form-control rounded-4" id="floatingInput" placeholder="name@example.com" v-model="email">
              <label for="floatingInput">Email address</label>
            </div>
            <div class="form-floating mb-3">
              <input type="password" class="form-control rounded-4" id="floatingPassword" placeholder="Password" v-model="password">
              <label for="floatingPassword">Password</label>
            </div>
            <button class="w-100 mb-2 btn btn-lg rounded-4 btn-primary" type="button" @click="submitform">Sign In</button>
          </form>
        </div>
      </div>
    </div>
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
          if (res.data.message === 'success') {
            window.location.replace('https://onebounce.me/panel')
          }
        })
    }
  }
}
</script>

<style>
.rounded-4 { border-radius: .5rem; }
.rounded-5 { border-radius: .75rem; }
.rounded-6 { border-radius: 1rem; }

.modal-sheet .modal-dialog {
  width: 380px;
  transition: bottom .75s ease-in-out;
}
.modal-sheet .modal-footer {
  padding-bottom: 2rem;
}

.modal-alert .modal-dialog {
  width: 380px;
}

.border-right { border-right: 1px solid #eee; }

.modal-tour .modal-dialog {
  width: 380px;
}
</style>
