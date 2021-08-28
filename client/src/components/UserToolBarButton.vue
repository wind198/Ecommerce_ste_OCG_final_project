<template>
  <button :class="buttonClass[0]">
    <i :class="buttonClass.slice(1)"></i>
  </button>
</template>

<script>
export default {
  name: "UserToolBarButton",
  props: {
    buttonClass: Array
  }
}
</script>

<style>
</style>