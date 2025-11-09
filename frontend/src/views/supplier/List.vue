<script setup lang="ts">
import type {Form} from '@/components/editor/SupplierEditor.vue';
import SupplierEditor from '@/components/editor/SupplierEditor.vue';
import Tag from '@/components/Tag.vue';
import { request } from '@/utils/axios';
import type { Supplier } from '@/utils/tables';
import { Plus, Search } from '@element-plus/icons-vue';
import { useUrlSearchParams } from '@vueuse/core';
import { reactive, ref, watchEffect } from 'vue';

const suppliers = ref<{
  data: Supplier[],
  total: number
} | null>(null)

const params = useUrlSearchParams<Record<string, string>>('history')
watchEffect(async () => suppliers.value = await request('GET', '/suppliers', params))

const isDialogOpen = ref(false)
const form = reactive<Form>({
  name: '', address: '', profile: '',
})

</script>

<template>

  <div class="p-4 flex flex-col gap-4">

    <div class="flex justify-between">
      <div class="text-xl">
        {{$t('supplier')}}
      </div>
      <el-button :icon="Plus" @click="isDialogOpen = true">
        {{$t('supplierList.addButton')}}
      </el-button>
    </div>

    <div class="flex gap-4">
      <el-input v-model="params.name" class="!w-64" :prefix-icon="Search" :placeholder="$t('search', { v: $t('supplier') })">
      </el-input>
    </div>

    <div v-if="suppliers" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div v-for="s in suppliers.data" :key="s.id" @click="$router.push(`/suppliers/${s.id}`)" class="card cursor-pointer flex flex-col gap-2">
        <div class="space-x-2 line-clamp-1">
          <span class="text-lg">{{s.name}}</span>
        </div>
        <div class="text-subtle">{{s.address}}</div>
        <div class="flex flex-wrap gap-1">
          <tag v-for="p in s.purchases" :key="p.id" :color="p.status.color" size="sm">
            {{p.name}}
          </tag>
        </div>
        <div class="mt-auto flex gap-2">
          <div class="text-subtle">{{$t('count', { k: $t('contact'), v: s.contactCount })}}</div>
          <div class="text-subtle">{{$t('count', { k: $t('purchase'), v: s.purchaseCount })}}</div>
        </div>
      </div>
    </div>

    <el-pagination v-if="suppliers" class="card" layout="prev, pager, next, total" :total="suppliers.total"
      :current-page="Number(params.page || 1)" :page-size="Number(params.size || 9)"
      @update:current-page="v => params.page = String(v)" @update:page-size="v => params.size = String(v)"
    ></el-pagination>
  </div>

  <el-dialog v-model="isDialogOpen" :title="$t('supplier.list.dialog.title')">
    <supplier-editor v-model="form">
    </supplier-editor>
    <template #footer>
      <el-button type="primary">
        {{$t('confirm')}}
      </el-button>
    </template>
  </el-dialog>

</template>
