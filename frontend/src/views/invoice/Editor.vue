<script setup lang="ts">

const isInvoiceDialogOpen = ref(false)
const invoiceForm = reactive<InvoiceForm>({
  name: '', number: '', amount: 0, issueDate: new Date(), dueDate: new Date(), statusId: '', profile: '',
})

</script>

<template>

</template>
