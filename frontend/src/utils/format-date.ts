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

export function formatDate(inputTime: string | number | Date): string {
  const now = dayjs();
  const time = dayjs(inputTime);

  const diffInMinutes = now.diff(time, 'minute');
  const diffInDays = now.diff(time, 'day');

  if (diffInMinutes < 1) {
    return '刚刚';
  } else if (diffInMinutes < 60) {
    return `${diffInMinutes}分钟前`;
  } else if (time.isToday()) {
    return time.format('HH:mm');
  } else if (time.isYesterday()) {
    return `昨天 ${time.format('HH:mm')}`;
  } else if (diffInDays < 7) {
    const weekDays = ['星期日','星期一','星期二','星期三','星期四','星期五','星期六'];
    return `${weekDays[time.day()]} ${time.format('HH:mm')}`;
  } else if (now.year() === time.year()) {
    return time.format('MM月DD日 HH:mm');
  } else {
    return time.format('YYYY年MM月DD日 HH:mm');
  }
}

