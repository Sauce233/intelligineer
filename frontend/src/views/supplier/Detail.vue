<script setup lang="ts">
import {request} from '@/utils/axios';
import type {Purchase, Supplier} from '@/utils/tables';
import {formatDate} from '@/utils/utils';
import {Back} from '@element-plus/icons-vue';
import {useRouteParams} from '@vueuse/router';
import {ref, watchEffect} from 'vue';

const id = useRouteParams('id')

const supplier = ref<Supplier | null>(null)

request<Supplier>('GET', `/suppliers/${id.value}`).then(res => supplier.value = res)

const purchases = ref<{
  data: Purchase[],
  total: number,
} | null>(null)

watchEffect(async () => purchases.value = await request('GET', '/purchases', { supplierId: id.value }))
</script>

<template>
  <div v-if="supplier" class="max-w-5xl mx-auto flex flex-col gap-4 p-4">

    <div class="card flex gap-2 items-center flex-wrap">
      <el-button @click="$router.back" :icon="Back">{{$t('back')}}</el-button>
      <div>{{supplier.name}}</div>
    </div>

    <div class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div
        v-for="v in [
          ['name', supplier.name, ''],
          ['createdAt', formatDate(supplier.createdAt), ''],
          ['updatedAt', formatDate(supplier.updatedAt), ''],
          ['address', supplier.address, 'col-span-2'],
        ]"
        :key="v[0]"
        :class="['flex items-center gap-2', v[2]]"
      >
        <div class="text-subtle">{{$t(v[0])}}</div>
        <div>{{v[1]}}</div>
      </div>
    </div>

    <div class="text-xl">
      {{$t('profile')}}
    </div>

    <div class="bg-subtle p-4 rounded-2xl">
      {{supplier.profile}}
    </div>

    <div class="p-2 flex gap-1">
      <router-link
        v-for="v in [
          [`/suppliers/${id}`, $t('contact')],
          [`/suppliers/${id}/purchases`, $t('purchase')],
        ]"
        :key="v[0]"
        :to="v[0]"
        class="p-2 rounded-2xl"
        exact-active-class="font-bold bg-subtle"
      >
        {{v[1]}}
      </router-link>
    </div>

    <router-view />

  </div>
</template>
