<script setup lang="ts">
import {request} from '@/utils/axios';
import {reactive} from 'vue';

const clientForm = reactive({
  name: '', address: '', profile: '',
})

</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">

    <el-form :model="clientForm">
      <el-form-item :label="$t('name')">
        <el-input v-model="clientForm.name" class="!w-64" />
      </el-form-item>
      <el-form-item :label="$t('address')">
        <el-input v-model="clientForm.address" type="textarea" />
      </el-form-item>
      <el-form-item :label="$t('profile')">
        <el-input v-model="clientForm.profile" type="textarea" />
      </el-form-item>
    </el-form>

    <el-button type="primary" @click="request('POST', '/clients', clientForm).then(id => id && $router.push(`/clients/${id}`))">
      {{$t('confirm')}}
    </el-button>

  </div>
</template>
