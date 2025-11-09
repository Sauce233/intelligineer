<script setup lang="ts">
import InvoiceEditor, {type Form} from '@/components/editor/InvoiceEditor.vue';
import Tag from '@/components/Tag.vue';
import { request } from '@/utils/axios';
import type { Invoice, InvoiceStatus } from '@/utils/tables';
import { formatDate, type List } from '@/utils/utils';
import { Plus, Search } from '@element-plus/icons-vue';
import { useUrlSearchParams } from '@vueuse/core';
import {useRouteParams} from '@vueuse/router';
import { reactive, ref, watchEffect } from 'vue';
import {useI18n} from 'vue-i18n';

const invoices = ref<List<Invoice> | null>(null)

const id = useRouteParams('id')
const params = useUrlSearchParams<Record<string, string>>('history')
watchEffect(async () => invoices.value = await request('GET', id.value ? `/projects/${id.value}/invoices` : '/invoices', params))

const invoiceStatuses = ref<InvoiceStatus[]>([])
request<InvoiceStatus[]>('GET', '/invoice-statuses').then(res => invoiceStatuses.value = res)

const isDialogOpen = ref(false)

const { t } = useI18n({
  messages: {},
})

const form = reactive<Form>({
  name: '', number:'', profile: '', amount: 0, time: [new Date(), new Date()], statusId: 0,
})

</script>

<template>
  <div class="flex justify-between">
    <div class="text-xl">
      {{$t('invoiceList.title')}}
    </div>
    <el-button @click="isDialogOpen = true" :icon="Plus">
      {{$t('projectDetail.addInvoiceButton')}}
    </el-button>
  </div>
  <div class="flex gap-4">
    <el-input v-model="params.name" :prefix-icon="Search" :placeholder="$t('invoiceList.searchNamePlaceholder')">
    </el-input>
    <el-input v-model="params.number" :prefix-icon="Search" :placeholder="$t('invoiceList.searchNumberPlaceholder')">
    </el-input>
    <el-button @click="delete params.statusId">{{$t('invoiceList.statusAll')}}</el-button>
    <el-segmented v-model="params.statusId" :options="invoiceStatuses" :props="{ value: 'id' }">
      <template #default="scope">
        {{scope.item.name}} {{scope.item.invoiceCount}}
      </template>
    </el-segmented>
  </div>

  <div v-if="invoices" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
    <div v-for="i in invoices.data" :key="i.id" @click="$router.push(`/invoices/${i.id}`)" class="card cursor-pointer flex flex-col gap-2">
      <div class="space-x-2 line-clamp-1">
        <tag :color="i.status.color">
          {{i.status.name}}
        </tag>
        <span class="text-lg">{{i.name}}</span>
      </div>
      <div class="text-subtle flex justify-between">
        <span>#{{i.number}}</span>
        <span class="font-bold">¥ {{i.amount.toLocaleString()}}</span>
      </div>
      <div class="text-subtle link" @click.stop="$router.push(`/projects/${i.project.id}`)">{{i.project.name}}</div>
      <div class="mt-auto text-subtle">
        <div>{{$t('invoiceList.issueDate')}}: {{formatDate(i.issueDate)}}</div>
        <div>{{$t('invoiceList.dueDate')}}: {{formatDate(i.dueDate)}}</div>
      </div>
    </div>
  </div>

  <el-pagination v-if="invoices" class="card" layout="prev, pager, next, total" :total="invoices.total"
    :current-page="Number(params.page || 1)" :page-size="Number(params.size || 9)"
    @update:current-page="v => params.page = String(v)" @update:page-size="v => params.size = String(v)"
  ></el-pagination>

  <el-dialog v-model="isDialogOpen" :title="$t('dialogTitle')">
    <invoice-editor v-model="form">
    </invoice-editor>
    <template #footer>
      <el-button>
        {{$t('confirm')}}
      </el-button>
    </template>
  </el-dialog>
</template>
