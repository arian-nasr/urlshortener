<template>
  <div class="modal modal-signin position-static d-block py-5" tabindex="-1" role="dialog" id="modalSignin">
    <div class="modal-dialog" role="document">
      <div class="modal-content rounded-5 shadow">
        <div class="modal-header p-5 pb-4 border-bottom-0">
          <!-- <h5 class="modal-title">Modal title</h5> -->
          <h2 class="fw-bold mb-0">URL Shortener</h2>
        </div>
        <div v-if="notshortened" class="modal-body p-5 pt-0">
          <form class="">
            <div class="form-floating mb-3">
              <input type="text" class="form-control rounded-4" id="floatingInput" placeholder="Long URL" v-model="url">
              <label for="floatingInput">Long URL</label>
            </div>
            <button class="w-100 mb-2 btn btn-lg rounded-4 btn-primary" type="button" @click="shorten">Shorten</button>
            <div v-if="errorShow" class="alert alert-danger rounded-4" role="alert">
              {{ errorMsg }}
            </div>
          </form>
        </div>
        <div v-else class="modal-body p-5 pt-0">
          <form class="">
            <div class="input-group mb-3">
              <input type="text" class="form-control" placeholder="Recipient's username" aria-label="Recipient's username" aria-describedby="basic-addon2" v-model="shortenedURL" ref="shortened">
              <div class="input-group-append">
                <button class="btn btn-outline-secondary" type="button">Copy</button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      url: '',
      errorShow: false,
      errorMsg: '',
      notshortened: true,
      shortenedURL: ''
    }
  },
  methods: {
    shorten: function () {
      this.shortened = true
      this.shortenedURL = 'https://onebounce.me/google'
    },
    clipboard: function () {
      this.$refs.shortened.select()
      navigator.clipboard.writeText(this.shortenedURL)
    }
  }
}
</script>

<style>
.rounded-4 { border-radius: .5rem !important; }
.rounded-5 { border-radius: .75rem !important; }
.rounded-6 { border-radius: 1rem !important; }

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
