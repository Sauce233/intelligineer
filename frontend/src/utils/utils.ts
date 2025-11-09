import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import isToday from 'dayjs/plugin/isToday';
import isYesterday from 'dayjs/plugin/isYesterday';
import weekday from 'dayjs/plugin/weekday';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import 'dayjs/locale/zh-cn';

dayjs.extend(relativeTime);
dayjs.extend(isToday);
dayjs.extend(isYesterday);
dayjs.extend(weekday);
dayjs.extend(localizedFormat);
dayjs.locale('zh-cn');

export interface List<T> {
  data: T[];
  total: number;
}

export function formatDate(input: string): string {
  return dayjs(input).format('YY/MM/DD hh:mm')
}

/*
type RelativeKey =
  | 'secondsAgo' | 'minutesAgo' | 'hoursAgo' | 'daysAgo'
  | 'monthsAgo' | 'yearsAgo'
  | 'inSeconds' | 'inMinutes' | 'inHours' | 'inDays'
  | 'inMonths' | 'inYears';

export function formatDate(
  input: Date | string,
  lang: string = 'zh-TW',
): string {
  const locale = timeLocales[lang] ?? timeLocales.en;
  const date = typeof input === 'string' ? new Date(input) : input;
  const now = new Date();
  const diffMs = date.getTime() - now.getTime();
  const diffSec = Math.round(diffMs / 1000);
  const absSec = Math.abs(diffSec);
  const isFuture = diffSec > 0;

  const getUnit = (absSec: number): [number, RelativeKey | 'justNow'] => {
    if (absSec < 10) return [0, 'justNow'];
    if (absSec < 60) return [absSec, isFuture ? 'inSeconds' : 'secondsAgo'];
    const mins = absSec / 60;
    if (mins < 60) return [Math.round(mins), isFuture ? 'inMinutes' : 'minutesAgo'];
    const hours = mins / 60;
    if (hours < 24) return [Math.round(hours), isFuture ? 'inHours' : 'hoursAgo'];
    const days = hours / 24;
    if (days < 30) return [Math.round(days), isFuture ? 'inDays' : 'daysAgo'];
    const months = days / 30;
    if (months < 12) return [Math.round(months), isFuture ? 'inMonths' : 'monthsAgo'];
    const years = months / 12;
    return [Math.round(years), isFuture ? 'inYears' : 'yearsAgo'];
  };

  const [n, key] = getUnit(absSec);

  if (key === 'justNow') return locale.justNow;

  // ✅ 这里我们明确限定类型，避免 dateFormat 被错误识别
  const formatter = locale[key] as (n: number) => string;
  if (typeof formatter === 'function') return formatter(n);

  // fallback 到绝对时间格式
  return locale.dateFormat(date);
}


export const timeLocales = {
  'zh-CN': {
    justNow: '刚刚',
    secondsAgo: (n: number) => `${n}秒前`,
    minutesAgo: (n: number) => `${n}分钟前`,
    hoursAgo: (n: number) => `${n}小时前`,
    daysAgo: (n: number) => `${n}天前`,
    monthsAgo: (n: number) => `${n}个月前`,
    yearsAgo: (n: number) => `${n}年前`,

    inSeconds: (n: number) => `${n}秒后`,
    inMinutes: (n: number) => `${n}分钟后`,
    inHours: (n: number) => `${n}小时后`,
    inDays: (n: number) => `${n}天后`,
    inMonths: (n: number) => `${n}个月后`,
    inYears: (n: number) => `${n}年后`,

    dateFormat: (date: Date) =>
      `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`,
  },

  'zh-TW': {
    justNow: '剛剛',
    secondsAgo: (n: number) => `${n}秒前`,
    minutesAgo: (n: number) => `${n}分鐘前`,
    hoursAgo: (n: number) => `${n}小時前`,
    daysAgo: (n: number) => `${n}天前`,
    monthsAgo: (n: number) => `${n}個月前`,
    yearsAgo: (n: number) => `${n}年前`,

    inSeconds: (n: number) => `${n}秒後`,
    inMinutes: (n: number) => `${n}分鐘後`,
    inHours: (n: number) => `${n}小時後`,
    inDays: (n: number) => `${n}天後`,
    inMonths: (n: number) => `${n}個月後`,
    inYears: (n: number) => `${n}年後`,

    dateFormat: (date: Date) =>
      `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`,
  },

  en: {
    justNow: 'just now',
    secondsAgo: (n: number) => `${n} seconds ago`,
    minutesAgo: (n: number) => `${n} minutes ago`,
    hoursAgo: (n: number) => `${n} hours ago`,
    daysAgo: (n: number) => `${n} days ago`,
    monthsAgo: (n: number) => `${n} months ago`,
    yearsAgo: (n: number) => `${n} years ago`,

    inSeconds: (n: number) => `in ${n} seconds`,
    inMinutes: (n: number) => `in ${n} minutes`,
    inHours: (n: number) => `in ${n} hours`,
    inDays: (n: number) => `in ${n} days`,
    inMonths: (n: number) => `in ${n} months`,
    inYears: (n: number) => `in ${n} years`,

    dateFormat: (date: Date) =>
      date.toLocaleDateString('en-US', { year: 'numeric', month: 'short', day: 'numeric' }),
  },
} as const;

type TimeLang = keyof typeof timeLocales;

*/

export const shortcuts = [
  {
    text: '最近1小时',
    value: () => {
      const end = new Date()
      const start = new Date(end.getTime() - 60 * 60 * 1000)
      return [start, end]
    },
  },
  {
    text: '最近6小时',
    value: () => {
      const end = new Date()
      const start = new Date(end.getTime() - 6 * 60 * 60 * 1000)
      return [start, end]
    },
  },
  {
    text: '今天',
    value: () => {
      const start = new Date()
      start.setHours(0, 0, 0, 0)
      const end = new Date()
      end.setHours(23, 59, 59, 999)
      return [start, end]
    },
  },
  {
    text: '昨天',
    value: () => {
      const start = new Date()
      start.setDate(start.getDate() - 1)
      start.setHours(0, 0, 0, 0)
      const end = new Date()
      end.setDate(end.getDate() - 1)
      end.setHours(23, 59, 59, 999)
      return [start, end]
    },
  },
  {
    text: '最近24小时',
    value: () => {
      const end = new Date()
      const start = new Date(end.getTime() - 24 * 60 * 60 * 1000)
      return [start, end]
    },
  },
  {
    text: '最近7天',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setDate(start.getDate() - 7)
      return [start, end]
    },
  },
  {
    text: '本周',
    value: () => {
      const now = new Date()
      const day = now.getDay() || 7
      const start = new Date(now)
      start.setDate(now.getDate() - day + 1)
      start.setHours(0, 0, 0, 0)
      const end = new Date()
      end.setHours(23, 59, 59, 999)
      return [start, end]
    },
  },
  {
    text: '上周',
    value: () => {
      const now = new Date()
      const day = now.getDay() || 7
      const end = new Date(now)
      end.setDate(now.getDate() - day)
      end.setHours(23, 59, 59, 999)
      const start = new Date(end)
      start.setDate(end.getDate() - 6)
      start.setHours(0, 0, 0, 0)
      return [start, end]
    },
  },
  {
    text: '本月',
    value: () => {
      const start = new Date()
      start.setDate(1)
      start.setHours(0, 0, 0, 0)
      const end = new Date()
      end.setHours(23, 59, 59, 999)
      return [start, end]
    },
  },
  {
    text: '上个月',
    value: () => {
      const start = new Date()
      start.setMonth(start.getMonth() - 1, 1)
      start.setHours(0, 0, 0, 0)
      const end = new Date(start)
      end.setMonth(end.getMonth() + 1)
      end.setDate(0)
      end.setHours(23, 59, 59, 999)
      return [start, end]
    },
  },
]

type ColorKey =
  | 'red' | 'orange' | 'amber' | 'yellow' | 'lime' | 'green'
  | 'emerald' | 'teal' | 'cyan' | 'sky' | 'blue' | 'indigo'
  | 'violet' | 'purple' | 'fuchsia' | 'pink' | 'rose'
  | 'slate' | 'gray' | 'zinc' | 'neutral' | 'stone';

type ClassType = 'bg' | 'text' | 'border';

// Tuple format: [border, text, bg]
const styles: Record<ColorKey, [string, string, string]> = {
  red: ['border-red-500', 'text-red-700', 'bg-red-100'],
  orange: ['border-orange-500', 'text-orange-700', 'bg-orange-100'],
  amber: ['border-amber-500', 'text-amber-700', 'bg-amber-100'],
  yellow: ['border-yellow-500', 'text-yellow-700', 'bg-yellow-100'],
  lime: ['border-lime-500', 'text-lime-700', 'bg-lime-100'],
  green: ['border-green-500', 'text-green-700', 'bg-green-100'],
  emerald: ['border-emerald-500', 'text-emerald-700', 'bg-emerald-100'],
  teal: ['border-teal-500', 'text-teal-700', 'bg-teal-100'],
  cyan: ['border-cyan-500', 'text-cyan-700', 'bg-cyan-100'],
  sky: ['border-sky-500', 'text-sky-700', 'bg-sky-100'],
  blue: ['border-blue-500', 'text-blue-700', 'bg-blue-100'],
  indigo: ['border-indigo-500', 'text-indigo-700', 'bg-indigo-100'],
  violet: ['border-violet-500', 'text-violet-700', 'bg-violet-100'],
  purple: ['border-purple-500', 'text-purple-700', 'bg-purple-100'],
  fuchsia: ['border-fuchsia-500', 'text-fuchsia-700', 'bg-fuchsia-100'],
  pink: ['border-pink-500', 'text-pink-700', 'bg-pink-100'],
  rose: ['border-rose-500', 'text-rose-700', 'bg-rose-100'],
  slate: ['border-slate-500', 'text-slate-700', 'bg-slate-100'],
  gray: ['border-gray-500', 'text-gray-700', 'bg-gray-100'],
  zinc: ['border-zinc-500', 'text-zinc-700', 'bg-zinc-100'],
  neutral: ['border-neutral-500', 'text-neutral-700', 'bg-neutral-100'],
  stone: ['border-stone-500', 'text-stone-700', 'bg-stone-100'],
};

const typeIndex: Record<ClassType, number> = { border: 0, text: 1, bg: 2 };

/**
 * Returns a Tailwind class string for the given color + type.
 * Falls back to gray if color is unknown.
 */
export function getColorClass(type: ClassType, color: string): string {
  const colorKey = color as ColorKey;
  return styles[colorKey]?.[typeIndex[type]] ?? styles.gray[typeIndex[type]];
}

export const fileIconMap: Record<
  string,
  { icon: string; color: string }
> = {
  // --- Office Documents ---
  'pdf': { icon: 'mdi:file-pdf-box', color: 'text-red-500' },
  'doc': { icon: 'mdi:file-word-box', color: 'text-blue-600' },
  'docx': { icon: 'mdi:file-word-box', color: 'text-blue-600' },
  'xls': { icon: 'mdi:file-excel-box', color: 'text-green-600' },
  'xlsx': { icon: 'mdi:file-excel-box', color: 'text-green-600' },
  'ppt': { icon: 'mdi:file-powerpoint-box', color: 'text-orange-500' },
  'pptx': { icon: 'mdi:file-powerpoint-box', color: 'text-orange-500' },
  'txt': { icon: 'mdi:file-document-outline', color: 'text-gray-500' },
  'csv': { icon: 'mdi:file-delimited', color: 'text-green-500' },
  'rtf': { icon: 'mdi:file-document-outline', color: 'text-gray-500' },

  // --- Images (used for project diagrams, reports, etc.) ---
  'jpg': { icon: 'mdi:file-image', color: 'text-amber-500' },
  'jpeg': { icon: 'mdi:file-image', color: 'text-amber-500' },
  'png': { icon: 'mdi:file-image', color: 'text-amber-500' },
  'gif': { icon: 'mdi:file-image', color: 'text-amber-500' },
  'svg': { icon: 'mdi:file-image', color: 'text-amber-500' },
  'webp': { icon: 'mdi:file-image', color: 'text-amber-500' },

  // --- Compressed / Package Files ---
  'zip': { icon: 'mdi:folder-zip', color: 'text-yellow-600' },
  'rar': { icon: 'mdi:folder-zip', color: 'text-yellow-600' },
  '7z': { icon: 'mdi:folder-zip', color: 'text-yellow-600' },
  'tar': { icon: 'mdi:folder-zip', color: 'text-yellow-600' },

  // --- Project / Configuration Files ---
  'json': { icon: 'mdi:code-json', color: 'text-emerald-600' },
  'xml': { icon: 'mdi:xml', color: 'text-indigo-600' },
  'yml': { icon: 'mdi:code-braces', color: 'text-indigo-600' },
  'yaml': { icon: 'mdi:code-braces', color: 'text-indigo-600' },
  'md': { icon: 'mdi:language-markdown', color: 'text-slate-600' },

  // --- Miscellaneous ---
  'log': { icon: 'mdi:file-document-outline', color: 'text-gray-500' },
  'bak': { icon: 'mdi:archive', color: 'text-gray-400' },
};

