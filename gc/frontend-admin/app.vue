<template>
  <NuxtLayout>
    <ClientOnly>
      <NuxtPage />
    </ClientOnly>
  </NuxtLayout>
</template>
<script lang="ts" setup></script>
<style lang="scss">
html,
#wrap {
  min-height: 100%;
  min-height: 100vh;
}
</style>
