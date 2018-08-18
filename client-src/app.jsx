'use strict';

import React from 'react';
import ReactDOM from 'react-dom';

const data = JSON.parse(document.currentScript.dataset.params);

const ComposeForm = (props) => (
  <div className="compose-container">
    <div className="compose">
      <form id="activity-form" action="/activity" method="post">
        <textarea name="body" placeholder="例) 今何してますか？　忙しいですか？"></textarea>
        <p><input type="submit" value="ぼやく" /></p>
      </form>
    </div>
  </div>
);
const SizeConfig = (props) => (
  <p>
    size:
    <select onChange={function (e) { location.href='/?size='+e.target.value; } } defaultValue={props.size}>
      {props.sizes.map((s) => (<option key={s} value={s}>{s}</option>))}
    </select>
  </p>
);
const EmptyActivitiesNotice = (props) => (
  <li>まだ投稿がありません。キミだけのオリジナルぼやきを投稿してみよう！</li>
);
const Activity = (props) => (
  <li>
    <ul>
      <li className="image"><img src={'/static/' + props.Username + '.jpg'} width="64" height="64" /></li>
      <li className="username">{ props.Username }</li>
      <li className="body">{ props.Body }</li>
    </ul>
  </li>
);
const ActivityList = (props) => (
  <ul className="activity">
    {props.activities.length <= 0 ? <EmptyActivitiesNotice /> : props.activities.map((a) => (<Activity key={a.Id} {...a} />))}
  </ul>
);
const Pager = (props) => (
  <ul className="pager pure-paginator">
    <li><a id="prevPage" {...props.prev_pager} /></li>
    <li id="currentPage">{props.page} / {props.total_page}</li>
    <li><a id="nextPage" {...props.next_pager} /></li>
  </ul>
);

const logoutHandler = () => {
  const r = new XMLHttpRequest();
  r.open('POST', '/logout');
  r.send();
  location.href='/login';
  return false;
};

ReactDOM.render(
  <div>
    <div className="pure-menu pure-menu-fixed">
    <h1>VOYAtter <span className="note">{data.user.Username} さんとしてログインしています　<a href="/logout" onClick={logoutHandler}>[logout]</a></span></h1>
    </div>

    <div id="activities">
      <ComposeForm {...data} />

      <div className="list">
        <SizeConfig {...data} />
        <ActivityList {...data} />
        <div id="footer">
          <Pager {...data} />
        </div>
      </div>
    </div>
  </div>,
  document.getElementById('app')
);
