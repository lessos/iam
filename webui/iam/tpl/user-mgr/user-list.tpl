<style type="text/css">
.pagination {
  margin: 10px 0;
}
</style>

<div id="iam-usermgr-list-alert" class="hide" style="margin:20px 0;"></div>
<div class="iam-div-light" id="iam-usermgr-list"></div>
<div id="iam-usermgr-list-pager"></div>

<script id="iam-usermgr-list-tpl" type="text/html">
<table class="table table-hover valign-middle">
  <thead>
    <tr>
      <th>Username</th>
      <th>Display Name</th>
      <th>Type</th>
      <th>Email</th>
      <th>Roles</th>
      <th>Status</th>
      <th>Updated</th>
      <th></th>
    </tr>
  </thead>
  <tbody >
{[~it.items :v]}
<tr>
  <td class="iam-monofont">{[=v.name]}</td>
  <td>{[=v.display_name]}</td>
  <td>{[=v._type]}</td>
  <td>{[=v.email]}</td>
  <td>
    {[~v.roles :rv]}
    {[~it._roles.items :drv]}
    {[ if (drv.id == rv) { ]}
    <div>{[=drv.name]}</div>
    {[ } ]}
    {[~]}
    {[~]}
  </td>
  <td>
    {[~it._statusls :sv]}
    {[ if (v.status == sv.status) { ]}{[=sv.title]}{[ } ]}
    {[~]}
  </td>
  <td>{[=l4i.UnixMillisecondFormat(v.updated, "Y-m-d")]}</td>
  <td align="right">
    <button class="pure-button button-small"
      onclick="iamAccMgr.FundNew('{[=v.name]}')">
      {[=l4i.T("Recharge")]}
    </button>
	{[? v._type_group]}
    <button class="pure-button button-small"
      onclick="iamUserMgr.GroupSetForm('{[=v.name]}')">
      {[=l4i.T("Settings")]}
    </button>
	{[??]}
    <button class="pure-button button-small"
      onclick="iamUserMgr.UserSetForm('{[=v.name]}')">
      {[=l4i.T("Settings")]}
    </button>
	{[?]}
  </td>
</tr>
{[~]}
</tbody>
</table>
</script>

<script id="iam-usermgr-list-pager-tpl" type="text/html">
<ul class="pagination pagination-sm">
  {[ if (it.FirstPageNumber > 0) { ]}
  <li><a href="#{[=it.FirstPageNumber]}">First</a></li>
  {[ } ]}
  {[~it.RangePages :v]}
  <li {[ if (v == it.CurrentPageNumber) { ]}class="active"{[ } ]}><a href="#{[=v]}">{[=v]}</a></li>
  {[~]}
  {[ if (it.LastPageNumber > 0) { ]}
  <li><a href="#{[=it.LastPageNumber]}">Last</a></li>
  {[ } ]}
</ul>
</script>


<script type="text/html" id="iam-usermgr-list-optools">
<li>
  <form id="iam-usermgr-list-query" action="#" onsubmit="iamUserMgr.UserList()" class="form-inlines">
    <input id="iam_usermgr_list_qry_text" type="text"
      class="form-control iam-query-input" 
      placeholder="Press Enter to Search" 
      value="">
  </form>
</li>
<li class="iam-btn iam-btn-primary">
  <a href="#" onclick="iamUserMgr.UserSetForm()">
     {[=l4i.T("New %s", l4i.T("User"))]} 
  </a>
</li>
</script>
