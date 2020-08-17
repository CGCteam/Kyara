import React, { useState, useEffect } from 'react';
import '@/components/calendar-heatmap/styles.css';
import { initTheme } from '@/utils/theme';
import './index.less';
import List from './components/list';
import { Pagination } from 'antd';

const Home = (props: any) => {
  useEffect(() => {
    initTheme(true);
  }, []);
  const [data] = useState({
    id: 1,
    type: 'top',
    title: '这是文章标题',
    summary:
      '文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要ellipsis文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要',
    auth: 'uvdream',
    time: '2020 08-10',
    comment: '3',
    img:
      'https://gitee.com/Uvdream/images/raw/master/images/20200806150710.png',
    icon: '',
  });
  const [data1] = useState({
    id: 2,
    type: 'icon',
    title: '这是文章标题',
    summary:
      '文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要ellipsis文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要要文章摘要文章摘要',
    auth: 'uvdream',
    time: '2020年 8月10日',
    comment: '3',
    img:
      'https://gitee.com/Uvdream/images/raw/master/images/20200806150710.png',
    icon: 'uv-wine',
  });
  const [data2] = useState({
    id: 3,
    type: 'leftImg',
    title: '这是文章标题',
    summary:
      '文文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要章摘要',
    auth: 'uvdream',
    time: '2020 08-10',
    comment: '3',
    img:
      'https://gitee.com/Uvdream/images/raw/master/images/20200806150710.png',
    icon: '',
  });
  const [data3] = useState({
    id: 4,
    type: 'normal',
    title: '这是文章标题',
    summary:
      '文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要',
    auth: 'uvdream',
    time: '2020 08-10',
    comment: '3',
    img:
      'https://gitee.com/Uvdream/images/raw/master/images/20200806150710.png',
    icon: '',
  });
  const [data4] = useState({
    id: 5,
    type: 'topImg',
    title: '这是文章标题',
    summary:
      '文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要文章摘要',
    auth: 'uvdream',
    time: '2020 08-10',
    comment: '3',
    img:
      'https://gitee.com/Uvdream/images/raw/master/images/20200806150710.png',
    icon: '',
  });
  const themeChange = (data: Boolean) => {
    initTheme(data);
  };
  return (
    <div className="home">
      <List data={data} />
      <List data={data1} />
      <List data={data2} />
      <List data={data3} />
      <List data={data4} />
      <div className="home-page">
        <section>
          <Pagination size="small" defaultCurrent={1} total={50} />
        </section>
      </div>
      {/* <button
        onClick={() => {
          themeChange(true);
        }}
      >
        切换
      </button>
      <button
        onClick={() => {
          themeChange(false);
        }}
      >
        切换
      </button>
      博客首页 */}
    </div>
  );
};
export default Home;