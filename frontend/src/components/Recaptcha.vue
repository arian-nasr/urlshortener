<template>
  <div id="g-recaptcha" class="g-recaptcha" :data-sitekey="sitekey">
  </div>
</template>

<script>
export default {
  data () {
    return {
      sitekey: '6LcKKGocAAAAABC0M9lYG563j_8ct4gB0kJ9xHhg',
      widgetId: 0
    }
  },
  methods: {
    execute () {
      window.grecaptcha.execute(this.widgetId)
    },
    reset () {
      window.grecaptcha.reset(this.widgetId)
    },
    render () {
      if (window.grecaptcha) {
        this.widgetId = window.grecaptcha.render('g-recaptcha', {
          sitekey: this.sitekey,
          size: 'invisible',
          // the callback executed when the user solve the recaptcha
          callback: (response) => {
            // emit an event called verify with the response as payload
            this.$emit('verify', response)
            // reset the recaptcha widget so you can execute it again
            this.reset()
          }
        })
      }
    }
  },
  mounted () {
    let recaptchaScript = document.createElement('script')
    recaptchaScript.setAttribute('src', 'https://www.google.com/recaptcha/api.js')
    document.head.appendChild(recaptchaScript)
    // render the recaptcha widget when the component is mounted
    this.render()
  }
}
</script>
