<template>
    <div class="login-page">
        <div class="form">
            <form class="login-form">
                <input type="text" placeholder="username"/>
                <input type="password" placeholder="password"/>
                <button>login</button>
            </form>
        </div>
    </div>
</template>

<script>
export default {
  created () {
    document.title = 'Login'
  }
}
</script>
