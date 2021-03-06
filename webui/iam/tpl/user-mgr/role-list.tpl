<style type="text/css">
.pagination {
  margin: 10px 0;
}
</style>

<div class="iam-div-light">
  <table class="table table-hover valign-middle">
    <thead>
      <tr>
        <th>ID</th>
        <th>Name</th>
        <th>Description</th>
        <th>Status</th>
        <th>Created</th>
        <th>Updated</th>
        <th></th>
      </tr>
    </thead>
    <tbody id="iam-usermgr-rolelist"></tbody>
  </table>
</div>

<script id="iam-usermgr-rolelist-tpl" type="text/html">
{[~it.items :v]}
<tr>
  <td class="iam-monofont">{[=v.id]}</td>
  <td>{[=v.name]}</td>
  <td>{[=v.desc]}</td>
  <td>
    {[~it._statusls :sv]}
    {[ if (v.status == sv.status) { ]}{[=sv.title]}{[ } ]}
    {[~]}
  </td>
  <td>{[=l4i.UnixMillisecondFormat(v.created, "Y-m-d")]}</td>
  <td>{[=l4i.UnixMillisecondFormat(v.updated, "Y-m-d")]}</td>
  <td align="right">
    <button class="pure-button button-small"
      onclick="iamUserMgr.RoleSet('{[=v.name]}')">
	  <span class="fa fa-cog"></span>
      {[=l4i.T("Settings")]}
    </button>
  </td>
</tr>
{[~]}
</script>

<script type="text/html" id="iam-usermgr-rolelist-optools">
<li class="iam-btn iam-btn-primary">
  <a href="#" onclick="iamUserMgr.RoleSet()">
    {[=l4i.T("New %s", l4i.T("Role"))]}
  </a>
</li>
</script>

