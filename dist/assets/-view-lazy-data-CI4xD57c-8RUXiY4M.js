function e(t){switch(t){case"backend":return`direction: down

Ui: {
  label: "UI"
}
Backend: {
  label: "Backend"

  Gateway: {
    label: "Gateway"
  }
  Users: {
    label: "Users"
  }
  Stats: {
    label: "Stats"
  }
  Content: {
    label: "Content"
  }
  Users_db: {
    label: "Users Database"
    shape: stored_data
  }
  Stats_db: {
    label: "Stats Databse"
    shape: stored_data
  }
  Content_db: {
    label: "Content Database"
    shape: stored_data
  }
}

Ui -> Backend.Gateway: "Use REST API"
Backend.Gateway -> Backend.Users: "[...]"
Backend.Gateway -> Backend.Stats: "[...]"
Backend.Gateway -> Backend.Content: "Get or post content"
Backend.Users -> Backend.Users_db: "Get or update profile data"
Backend.Stats -> Backend.Stats_db: "Store and calculate stats"
Backend.Content -> Backend.Content_db: "Store content"
`;case"index":return`direction: down

User: {
  label: "User"
  shape: person
}
Ui: {
  label: "UI"
}
Backend: {
  label: "Backend"
}

User -> Ui: "Use graphical interface"
Ui -> Backend: "Use REST API"
`;default:throw new Error("Unknown viewId: "+t)}}function n(t){switch(t){case"backend":return`digraph {
    graph [TBbalance=min,
        bgcolor=transparent,
        compound=true,
        fontname=Arial,
        fontsize=15,
        labeljust=l,
        labelloc=t,
        layout=dot,
        nodesep=1.528,
        outputorder=nodesfirst,
        pad=0.209,
        rankdir=TB,
        ranksep=1.667,
        splines=spline
    ];
    node [color="#2563eb",
        fillcolor="#3b82f6",
        fontcolor="#eff6ff",
        fontname=Arial,
        label="\\N",
        penwidth=0,
        shape=rect,
        style=filled
    ];
    edge [arrowsize=0.75,
        color="#6E6E6E",
        fontcolor="#C6C6C6",
        fontname=Arial,
        fontsize=14,
        penwidth=2
    ];
    subgraph cluster_backend {
        graph [color="#1b3d88",
            fillcolor="#194b9e",
            label=<<FONT POINT-SIZE="11" COLOR="#bfdbfeb3"><B>BACKEND</B></FONT>>,
            likec4_depth=1,
            likec4_id=backend,
            likec4_level=0,
            margin=40,
            style=filled
        ];
        gateway [group=backend,
            height=2.5,
            label=<<TABLE BORDER="0" CELLPADDING="0" CELLSPACING="4"><TR><TD ROWSPAN="2" WIDTH="76"> </TD><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="19">Gateway</FONT></TD><TD ROWSPAN="2" WIDTH="16"> </TD></TR><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="15" COLOR="#bfdbfe">Exposes backend to the outside</FONT></TD></TR></TABLE>>,
            likec4_id="backend.gateway",
            likec4_level=1,
            margin="0.112,0.223",
            width=4.445];
        users [group=backend,
            height=2.5,
            label=<<TABLE BORDER="0" CELLPADDING="0" CELLSPACING="4"><TR><TD ROWSPAN="2" WIDTH="76"> </TD><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="19">Users</FONT></TD><TD ROWSPAN="2" WIDTH="16"> </TD></TR><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="15" COLOR="#bfdbfe">Responsible for user profiles</FONT></TD></TR></TABLE>>,
            likec4_id="backend.users",
            likec4_level=1,
            margin="0.112,0.223",
            width=4.445];
        stats [group=backend,
            height=2.5,
            label=<<TABLE BORDER="0" CELLPADDING="0" CELLSPACING="4"><TR><TD ROWSPAN="2" WIDTH="76"> </TD><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="19">Stats</FONT></TD><TD ROWSPAN="2" WIDTH="16"> </TD></TR><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="15" COLOR="#bfdbfe">Accumulates and manages stats</FONT></TD></TR></TABLE>>,
            likec4_id="backend.stats",
            likec4_level=1,
            margin="0.112,0.223",
            width=4.445];
        content [group=backend,
            height=2.5,
            label=<<TABLE BORDER="0" CELLPADDING="0" CELLSPACING="4"><TR><TD ROWSPAN="2" WIDTH="76"> </TD><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="19">Content</FONT></TD><TD ROWSPAN="2" WIDTH="16"> </TD></TR><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="15" COLOR="#bfdbfe">Responsible for all user-generated<BR/>content</FONT></TD></TR></TABLE>>,
            likec4_id="backend.content",
            likec4_level=1,
            margin="0.112,0.223",
            width=4.445];
        users_db [group=backend,
            height=2.5,
            label=<<TABLE BORDER="0" CELLPADDING="0" CELLSPACING="4"><TR><TD ROWSPAN="2" WIDTH="76"> </TD><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="19">Users Database</FONT></TD><TD ROWSPAN="2" WIDTH="16"> </TD></TR><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="15" COLOR="#bfdbfe">Stores profile data</FONT></TD></TR></TABLE>>,
            likec4_id="backend.users_db",
            likec4_level=1,
            margin="0.112,0",
            penwidth=2,
            shape=cylinder,
            width=4.445];
        stats_db [group=backend,
            height=2.5,
            label=<<TABLE BORDER="0" CELLPADDING="0" CELLSPACING="4"><TR><TD ROWSPAN="2" WIDTH="76"> </TD><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="19">Stats Databse</FONT></TD><TD ROWSPAN="2" WIDTH="16"> </TD></TR><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="15" COLOR="#bfdbfe">Stores and processes stats</FONT></TD></TR></TABLE>>,
            likec4_id="backend.stats_db",
            likec4_level=1,
            margin="0.112,0",
            penwidth=2,
            shape=cylinder,
            width=4.445];
        content_db [group=backend,
            height=2.5,
            label=<<TABLE BORDER="0" CELLPADDING="0" CELLSPACING="4"><TR><TD ROWSPAN="2" WIDTH="76"> </TD><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="19">Content Database</FONT></TD><TD ROWSPAN="2" WIDTH="16"> </TD></TR><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="15" COLOR="#bfdbfe">Stores user-generated content</FONT></TD></TR></TABLE>>,
            likec4_id="backend.content_db",
            likec4_level=1,
            margin="0.112,0",
            penwidth=2,
            shape=cylinder,
            width=4.445];
    }
    ui [height=2.5,
        label=<<TABLE BORDER="0" CELLPADDING="0" CELLSPACING="4"><TR><TD ROWSPAN="2" WIDTH="76"> </TD><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="19">UI</FONT></TD><TD ROWSPAN="2" WIDTH="16"> </TD></TR><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="15" COLOR="#bfdbfe">Frontend in general</FONT></TD></TR></TABLE>>,
        likec4_id=ui,
        likec4_level=0,
        margin="0.112,0.306",
        width=4.445];
    ui -> gateway [label=<<TABLE BORDER="0" CELLPADDING="3" CELLSPACING="0" BGCOLOR="#18191bA0"><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="14">Use REST API</FONT></TD></TR></TABLE>>,
        likec4_id="1b9rrvo",
        minlen=1,
        style=dashed];
    gateway -> users [label=<<TABLE BORDER="0" CELLPADDING="3" CELLSPACING="0" BGCOLOR="#18191bA0"><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="14"><B>[...]</B></FONT></TD></TR></TABLE>>,
        likec4_id="1y7yvtc",
        style=dashed,
        weight=2];
    gateway -> stats [label=<<TABLE BORDER="0" CELLPADDING="3" CELLSPACING="0" BGCOLOR="#18191bA0"><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="14"><B>[...]</B></FONT></TD></TR></TABLE>>,
        likec4_id="1y6nwc3",
        style=dashed,
        weight=2];
    gateway -> content [label=<<TABLE BORDER="0" CELLPADDING="3" CELLSPACING="0" BGCOLOR="#18191bA0"><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="14">Get or post content</FONT></TD></TR></TABLE>>,
        likec4_id="1hxso0r",
        style=dashed,
        weight=2];
    users -> users_db [label=<<TABLE BORDER="0" CELLPADDING="3" CELLSPACING="0" BGCOLOR="#18191bA0"><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="14">Get or update profile data</FONT></TD></TR></TABLE>>,
        likec4_id=xii4pf,
        minlen=1,
        style=dashed];
    stats -> stats_db [label=<<TABLE BORDER="0" CELLPADDING="3" CELLSPACING="0" BGCOLOR="#18191bA0"><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="14">Store and calculate stats</FONT></TD></TR></TABLE>>,
        likec4_id="1id4jrn",
        minlen=1,
        style=dashed];
    content -> content_db [label=<<TABLE BORDER="0" CELLPADDING="3" CELLSPACING="0" BGCOLOR="#18191bA0"><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="14">Store content</FONT></TD></TR></TABLE>>,
        likec4_id=wbrnrn,
        minlen=1,
        style=dashed];
}
`;case"index":return`digraph {
    graph [TBbalance=min,
        bgcolor=transparent,
        compound=true,
        fontname=Arial,
        fontsize=15,
        labeljust=l,
        labelloc=t,
        layout=dot,
        nodesep=1.528,
        outputorder=nodesfirst,
        pad=0.209,
        rankdir=TB,
        ranksep=1.667,
        splines=spline
    ];
    node [color="#2563eb",
        fillcolor="#3b82f6",
        fontcolor="#eff6ff",
        fontname=Arial,
        label="\\N",
        penwidth=0,
        shape=rect,
        style=filled
    ];
    edge [arrowsize=0.75,
        color="#6E6E6E",
        fontcolor="#C6C6C6",
        fontname=Arial,
        fontsize=14,
        penwidth=2
    ];
    user [height=2.5,
        label=<<TABLE BORDER="0" CELLPADDING="0" CELLSPACING="4"><TR><TD><FONT POINT-SIZE="19">User</FONT></TD></TR><TR><TD><FONT POINT-SIZE="15" COLOR="#bfdbfe">End user of the platform</FONT></TD></TR></TABLE>>,
        likec4_id=user,
        likec4_level=0,
        margin="0.223,0.223",
        width=4.445];
    ui [height=2.5,
        label=<<TABLE BORDER="0" CELLPADDING="0" CELLSPACING="4"><TR><TD ROWSPAN="2" WIDTH="76"> </TD><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="19">UI</FONT></TD><TD ROWSPAN="2" WIDTH="16"> </TD></TR><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="15" COLOR="#bfdbfe">Frontend in general</FONT></TD></TR></TABLE>>,
        likec4_id=ui,
        likec4_level=0,
        margin="0.112,0.306",
        width=4.445];
    user -> ui [label=<<TABLE BORDER="0" CELLPADDING="3" CELLSPACING="0" BGCOLOR="#18191bA0"><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="14">Use graphical interface</FONT></TD></TR></TABLE>>,
        likec4_id="1uutevb",
        minlen=1,
        style=dashed];
    backend [height=2.5,
        label=<<TABLE BORDER="0" CELLPADDING="0" CELLSPACING="4"><TR><TD ROWSPAN="2" WIDTH="76"> </TD><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="19">Backend</FONT></TD><TD ROWSPAN="2" WIDTH="16"> </TD></TR><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="15" COLOR="#bfdbfe">Consists of microservices. This is<BR/>the part we are working on</FONT></TD></TR></TABLE>>,
        likec4_id=backend,
        likec4_level=0,
        margin="0.112,0.223",
        width=4.445];
    ui -> backend [label=<<TABLE BORDER="0" CELLPADDING="3" CELLSPACING="0" BGCOLOR="#18191bA0"><TR><TD ALIGN="TEXT" BALIGN="LEFT"><FONT POINT-SIZE="14">Use REST API</FONT></TD></TR></TABLE>>,
        likec4_id=tj8vma,
        minlen=1,
        style=dashed];
}
`;default:throw new Error("Unknown viewId: "+t)}}function a(t){switch(t){case"backend":return`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN"
 "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<!-- Generated by graphviz version 12.2.1 (0)
 -->
<!-- Pages: 1 -->
<svg width="1360pt" height="1236pt"
 viewBox="0.00 0.00 1360.10 1235.70" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
<g id="graph0" class="graph" transform="scale(1 1) rotate(0) translate(15.05 1220.65)">
<g id="clust1" class="cluster">
<title>cluster_backend</title>
<polygon fill="#194b9e" stroke="#1b3d88" points="8,-8 8,-934.8 1322,-934.8 1322,-8 8,-8"/>
<text text-anchor="start" x="16" y="-921.9" font-family="Arial" font-weight="bold" font-size="11.00" fill="#bfdbfe" fill-opacity="0.701961">BACKEND</text>
</g>
<!-- gateway -->
<g id="node1" class="node">
<title>gateway</title>
<polygon fill="#3b82f6" stroke="#2563eb" stroke-width="0" points="816.39,-873.6 479.61,-873.6 479.61,-693.6 816.39,-693.6 816.39,-873.6"/>
<text text-anchor="start" x="527.73" y="-779.4" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="571.68" y="-788.9" font-family="Arial" font-size="19.00" fill="#eff6ff">Gateway</text>
<text text-anchor="start" x="794.38" y="-779.4" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="571.68" y="-765.7" font-family="Arial" font-size="15.00" fill="#bfdbfe">Exposes backend to the outside</text>
</g>
<!-- users -->
<g id="node2" class="node">
<title>users</title>
<polygon fill="#3b82f6" stroke="#2563eb" stroke-width="0" points="368.02,-550.8 47.98,-550.8 47.98,-370.8 368.02,-370.8 368.02,-550.8"/>
<text text-anchor="start" x="99.01" y="-456.6" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="142.95" y="-466.1" font-family="Arial" font-size="19.00" fill="#eff6ff">Users</text>
<text text-anchor="start" x="343.1" y="-456.6" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="142.95" y="-442.9" font-family="Arial" font-size="15.00" fill="#bfdbfe">Responsible for user profiles</text>
</g>
<!-- stats -->
<g id="node3" class="node">
<title>stats</title>
<polygon fill="#3b82f6" stroke="#2563eb" stroke-width="0" points="818.46,-550.8 477.54,-550.8 477.54,-370.8 818.46,-370.8 818.46,-550.8"/>
<text text-anchor="start" x="525.66" y="-456.6" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="569.61" y="-466.1" font-family="Arial" font-size="19.00" fill="#eff6ff">Stats</text>
<text text-anchor="start" x="796.45" y="-456.6" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="569.61" y="-442.9" font-family="Arial" font-size="15.00" fill="#bfdbfe">Accumulates and manages stats</text>
</g>
<!-- content -->
<g id="node4" class="node">
<title>content</title>
<polygon fill="#3b82f6" stroke="#2563eb" stroke-width="0" points="1281.71,-550.8 928.29,-550.8 928.29,-370.8 1281.71,-370.8 1281.71,-550.8"/>
<text text-anchor="start" x="976.41" y="-456.6" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="1020.35" y="-475.1" font-family="Arial" font-size="19.00" fill="#eff6ff">Content</text>
<text text-anchor="start" x="1259.7" y="-456.6" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="1020.35" y="-451.9" font-family="Arial" font-size="15.00" fill="#bfdbfe">Responsible for all user&#45;generated</text>
<text text-anchor="start" x="1020.35" y="-433.9" font-family="Arial" font-size="15.00" fill="#bfdbfe">content</text>
</g>
<!-- users_db -->
<g id="node5" class="node">
<title>users_db</title>
<path fill="#3b82f6" stroke="#2563eb" stroke-width="2" d="M368.02,-211.64C368.02,-220.67 296.3,-228 208,-228 119.7,-228 47.98,-220.67 47.98,-211.64 47.98,-211.64 47.98,-64.36 47.98,-64.36 47.98,-55.33 119.7,-48 208,-48 296.3,-48 368.02,-55.33 368.02,-64.36 368.02,-64.36 368.02,-211.64 368.02,-211.64"/>
<path fill="none" stroke="#2563eb" stroke-width="2" d="M368.02,-211.64C368.02,-202.61 296.3,-195.27 208,-195.27 119.7,-195.27 47.98,-202.61 47.98,-211.64"/>
<text text-anchor="start" x="125.94" y="-133.8" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="169.89" y="-143.3" font-family="Arial" font-size="19.00" fill="#eff6ff">Users Database</text>
<text text-anchor="start" x="316.17" y="-133.8" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="169.89" y="-120.1" font-family="Arial" font-size="15.00" fill="#bfdbfe">Stores profile data</text>
</g>
<!-- stats_db -->
<g id="node6" class="node">
<title>stats_db</title>
<path fill="#3b82f6" stroke="#2563eb" stroke-width="2" d="M808.02,-211.64C808.02,-220.67 736.3,-228 648,-228 559.7,-228 487.98,-220.67 487.98,-211.64 487.98,-211.64 487.98,-64.36 487.98,-64.36 487.98,-55.33 559.7,-48 648,-48 736.3,-48 808.02,-55.33 808.02,-64.36 808.02,-64.36 808.02,-211.64 808.02,-211.64"/>
<path fill="none" stroke="#2563eb" stroke-width="2" d="M808.02,-211.64C808.02,-202.61 736.3,-195.27 648,-195.27 559.7,-195.27 487.98,-202.61 487.98,-211.64"/>
<text text-anchor="start" x="543.59" y="-133.8" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="587.54" y="-143.3" font-family="Arial" font-size="19.00" fill="#eff6ff">Stats Databse</text>
<text text-anchor="start" x="778.52" y="-133.8" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="587.54" y="-120.1" font-family="Arial" font-size="15.00" fill="#bfdbfe">Stores and processes stats</text>
</g>
<!-- content_db -->
<g id="node7" class="node">
<title>content_db</title>
<path fill="#3b82f6" stroke="#2563eb" stroke-width="2" d="M1268.38,-211.64C1268.38,-220.67 1195.15,-228 1105,-228 1014.85,-228 941.62,-220.67 941.62,-211.64 941.62,-211.64 941.62,-64.36 941.62,-64.36 941.62,-55.33 1014.85,-48 1105,-48 1195.15,-48 1268.38,-55.33 1268.38,-64.36 1268.38,-64.36 1268.38,-211.64 1268.38,-211.64"/>
<path fill="none" stroke="#2563eb" stroke-width="2" d="M1268.38,-211.64C1268.38,-202.61 1195.15,-195.27 1105,-195.27 1014.85,-195.27 941.62,-202.61 941.62,-211.64"/>
<text text-anchor="start" x="989.74" y="-133.8" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="1033.69" y="-143.3" font-family="Arial" font-size="19.00" fill="#eff6ff">Content Database</text>
<text text-anchor="start" x="1246.37" y="-133.8" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="1033.69" y="-120.1" font-family="Arial" font-size="15.00" fill="#bfdbfe">Stores user&#45;generated content</text>
</g>
<!-- ui -->
<g id="node8" class="node">
<title>ui</title>
<polygon fill="#3b82f6" stroke="#2563eb" stroke-width="0" points="808.02,-1205.6 487.98,-1205.6 487.98,-1025.6 808.02,-1025.6 808.02,-1205.6"/>
<text text-anchor="start" x="569.01" y="-1111.4" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="612.96" y="-1120.9" font-family="Arial" font-size="19.00" fill="#eff6ff">UI</text>
<text text-anchor="start" x="753.1" y="-1111.4" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="612.96" y="-1097.7" font-family="Arial" font-size="15.00" fill="#bfdbfe">Frontend in general</text>
</g>
<!-- gateway&#45;&gt;users -->
<g id="edge2" class="edge">
<title>gateway&#45;&gt;users</title>
<path fill="none" stroke="#6e6e6e" stroke-width="2" stroke-dasharray="5,2" d="M526.02,-693.67C467.6,-651.07 397.59,-600.03 338.15,-556.69"/>
<polygon fill="#6e6e6e" stroke="#6e6e6e" stroke-width="2" points="339.89,-554.71 332.28,-552.41 336.8,-558.95 339.89,-554.71"/>
<polygon fill="#18191b" fill-opacity="0.627451" stroke="none" points="441.19,-610.8 441.19,-633.6 468.19,-633.6 468.19,-610.8 441.19,-610.8"/>
<text text-anchor="start" x="444.19" y="-619" font-family="Arial" font-weight="bold" font-size="14.00" fill="#c6c6c6">[...]</text>
</g>
<!-- gateway&#45;&gt;stats -->
<g id="edge3" class="edge">
<title>gateway&#45;&gt;stats</title>
<path fill="none" stroke="#6e6e6e" stroke-width="2" stroke-dasharray="5,2" d="M648,-693.67C648,-652.47 648,-603.36 648,-560.97"/>
<polygon fill="#6e6e6e" stroke="#6e6e6e" stroke-width="2" points="650.63,-561.16 648,-553.66 645.38,-561.16 650.63,-561.16"/>
<polygon fill="#18191b" fill-opacity="0.627451" stroke="none" points="648,-610.8 648,-633.6 674.99,-633.6 674.99,-610.8 648,-610.8"/>
<text text-anchor="start" x="651" y="-619" font-family="Arial" font-weight="bold" font-size="14.00" fill="#c6c6c6">[...]</text>
</g>
<!-- gateway&#45;&gt;content -->
<g id="edge4" class="edge">
<title>gateway&#45;&gt;content</title>
<path fill="none" stroke="#6e6e6e" stroke-width="2" stroke-dasharray="5,2" d="M774.69,-693.67C835.49,-650.99 908.38,-599.82 970.2,-556.43"/>
<polygon fill="#6e6e6e" stroke="#6e6e6e" stroke-width="2" points="971.35,-558.83 975.98,-552.37 968.33,-554.53 971.35,-558.83"/>
<polygon fill="#18191b" fill-opacity="0.627451" stroke="none" points="890.2,-610.8 890.2,-633.6 1015.27,-633.6 1015.27,-610.8 890.2,-610.8"/>
<text text-anchor="start" x="893.2" y="-618" font-family="Arial" font-size="14.00" fill="#c6c6c6">Get or post content</text>
</g>
<!-- users&#45;&gt;users_db -->
<g id="edge5" class="edge">
<title>users&#45;&gt;users_db</title>
<path fill="none" stroke="#6e6e6e" stroke-width="2" stroke-dasharray="5,2" d="M208,-370.87C208,-330.01 208,-281.38 208,-239.23"/>
<polygon fill="#6e6e6e" stroke="#6e6e6e" stroke-width="2" points="210.63,-239.47 208,-231.97 205.38,-239.47 210.63,-239.47"/>
<polygon fill="#18191b" fill-opacity="0.627451" stroke="none" points="208,-288 208,-310.8 372.77,-310.8 372.77,-288 208,-288"/>
<text text-anchor="start" x="211" y="-295.2" font-family="Arial" font-size="14.00" fill="#c6c6c6">Get or update profile data</text>
</g>
<!-- stats&#45;&gt;stats_db -->
<g id="edge6" class="edge">
<title>stats&#45;&gt;stats_db</title>
<path fill="none" stroke="#6e6e6e" stroke-width="2" stroke-dasharray="5,2" d="M648,-370.87C648,-330.01 648,-281.38 648,-239.23"/>
<polygon fill="#6e6e6e" stroke="#6e6e6e" stroke-width="2" points="650.63,-239.47 648,-231.97 645.38,-239.47 650.63,-239.47"/>
<polygon fill="#18191b" fill-opacity="0.627451" stroke="none" points="648,-288 648,-310.8 807.31,-310.8 807.31,-288 648,-288"/>
<text text-anchor="start" x="651" y="-295.2" font-family="Arial" font-size="14.00" fill="#c6c6c6">Store and calculate stats</text>
</g>
<!-- content&#45;&gt;content_db -->
<g id="edge7" class="edge">
<title>content&#45;&gt;content_db</title>
<path fill="none" stroke="#6e6e6e" stroke-width="2" stroke-dasharray="5,2" d="M1105,-370.87C1105,-330.01 1105,-281.38 1105,-239.23"/>
<polygon fill="#6e6e6e" stroke="#6e6e6e" stroke-width="2" points="1107.63,-239.47 1105,-231.97 1102.38,-239.47 1107.63,-239.47"/>
<polygon fill="#18191b" fill-opacity="0.627451" stroke="none" points="1105,-288 1105,-310.8 1194.28,-310.8 1194.28,-288 1105,-288"/>
<text text-anchor="start" x="1108" y="-295.2" font-family="Arial" font-size="14.00" fill="#c6c6c6">Store content</text>
</g>
<!-- ui&#45;&gt;gateway -->
<g id="edge1" class="edge">
<title>ui&#45;&gt;gateway</title>
<path fill="none" stroke="#6e6e6e" stroke-width="2" stroke-dasharray="5,2" d="M648,-1025.73C648,-981.9 648,-928.88 648,-883.74"/>
<polygon fill="#6e6e6e" stroke="#6e6e6e" stroke-width="2" points="650.63,-883.87 648,-876.37 645.38,-883.87 650.63,-883.87"/>
<polygon fill="#18191b" fill-opacity="0.627451" stroke="none" points="648,-942.8 648,-965.6 746.58,-965.6 746.58,-942.8 648,-942.8"/>
<text text-anchor="start" x="651" y="-950" font-family="Arial" font-size="14.00" fill="#c6c6c6">Use REST API</text>
</g>
</g>
</svg>
`;case"index":return`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN"
 "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<!-- Generated by graphviz version 12.2.1 (0)
 -->
<!-- Pages: 1 -->
<svg width="376pt" height="856pt"
 viewBox="0.00 0.00 375.94 855.70" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
<g id="graph0" class="graph" transform="scale(1 1) rotate(0) translate(15.05 840.65)">
<!-- user -->
<g id="node1" class="node">
<title>user</title>
<polygon fill="#3b82f6" stroke="#2563eb" stroke-width="0" points="332.94,-825.6 12.9,-825.6 12.9,-645.6 332.94,-645.6 332.94,-825.6"/>
<text text-anchor="start" x="152.86" y="-740.9" font-family="Arial" font-size="19.00" fill="#eff6ff">User</text>
<text text-anchor="start" x="92.88" y="-717.7" font-family="Arial" font-size="15.00" fill="#bfdbfe">End user of the platform</text>
</g>
<!-- ui -->
<g id="node2" class="node">
<title>ui</title>
<polygon fill="#3b82f6" stroke="#2563eb" stroke-width="0" points="332.94,-502.8 12.9,-502.8 12.9,-322.8 332.94,-322.8 332.94,-502.8"/>
<text text-anchor="start" x="93.93" y="-408.6" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="137.88" y="-418.1" font-family="Arial" font-size="19.00" fill="#eff6ff">UI</text>
<text text-anchor="start" x="278.02" y="-408.6" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="137.88" y="-394.9" font-family="Arial" font-size="15.00" fill="#bfdbfe">Frontend in general</text>
</g>
<!-- backend -->
<g id="node3" class="node">
<title>backend</title>
<polygon fill="#3b82f6" stroke="#2563eb" stroke-width="0" points="345.84,-180 0,-180 0,0 345.84,0 345.84,-180"/>
<text text-anchor="start" x="48.12" y="-85.8" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="92.06" y="-104.3" font-family="Arial" font-size="19.00" fill="#eff6ff">Backend</text>
<text text-anchor="start" x="323.83" y="-85.8" font-family="Arial" font-size="14.00" fill="#eff6ff"> </text>
<text text-anchor="start" x="92.06" y="-81.1" font-family="Arial" font-size="15.00" fill="#bfdbfe">Consists of microservices. This is</text>
<text text-anchor="start" x="92.06" y="-63.1" font-family="Arial" font-size="15.00" fill="#bfdbfe">the part we are working on</text>
</g>
<!-- user&#45;&gt;ui -->
<g id="edge1" class="edge">
<title>user&#45;&gt;ui</title>
<path fill="none" stroke="#6e6e6e" stroke-width="2" stroke-dasharray="5,2" d="M172.92,-645.67C172.92,-604.47 172.92,-555.36 172.92,-512.97"/>
<polygon fill="#6e6e6e" stroke="#6e6e6e" stroke-width="2" points="175.54,-513.16 172.92,-505.66 170.29,-513.16 175.54,-513.16"/>
<polygon fill="#18191b" fill-opacity="0.627451" stroke="none" points="172.92,-562.8 172.92,-585.6 322.11,-585.6 322.11,-562.8 172.92,-562.8"/>
<text text-anchor="start" x="175.92" y="-570" font-family="Arial" font-size="14.00" fill="#c6c6c6">Use graphical interface</text>
</g>
<!-- ui&#45;&gt;backend -->
<g id="edge2" class="edge">
<title>ui&#45;&gt;backend</title>
<path fill="none" stroke="#6e6e6e" stroke-width="2" stroke-dasharray="5,2" d="M172.92,-322.87C172.92,-281.67 172.92,-232.56 172.92,-190.17"/>
<polygon fill="#6e6e6e" stroke="#6e6e6e" stroke-width="2" points="175.54,-190.36 172.92,-182.86 170.29,-190.36 175.54,-190.36"/>
<polygon fill="#18191b" fill-opacity="0.627451" stroke="none" points="172.92,-240 172.92,-262.8 271.5,-262.8 271.5,-240 172.92,-240"/>
<text text-anchor="start" x="175.92" y="-247.2" font-family="Arial" font-size="14.00" fill="#c6c6c6">Use REST API</text>
</g>
</g>
</svg>
`;default:throw new Error("Unknown viewId: "+t)}}function l(t){switch(t){case"backend":return`---
title: "Backend"
---
graph TB
  Ui[UI]
  subgraph Backend["Backend"]
    Backend.Gateway[Gateway]
    Backend.Users[Users]
    Backend.Stats[Stats]
    Backend.Content[Content]
    Backend.Users_db([Users Database])
    Backend.Stats_db([Stats Databse])
    Backend.Content_db([Content Database])
  end
  Ui -. "Use REST API" .-> Backend.Gateway
  Backend.Gateway -. "[...]" .-> Backend.Users
  Backend.Gateway -. "[...]" .-> Backend.Stats
  Backend.Gateway -. "Get or post content" .-> Backend.Content
  Backend.Users -. "Get or update profile data" .-> Backend.Users_db
  Backend.Stats -. "Store and calculate stats" .-> Backend.Stats_db
  Backend.Content -. "Store content" .-> Backend.Content_db
`;case"index":return`---
title: "System Landscape"
---
graph TB
  User[fa:fa-user User]
  Ui[UI]
  Backend[Backend]
  User -. "Use graphical interface" .-> Ui
  Ui -. "Use REST API" .-> Backend
`;default:throw new Error("Unknown viewId: "+t)}}export{e as d2Source,n as dotSource,l as mmdSource,a as svgSource};
