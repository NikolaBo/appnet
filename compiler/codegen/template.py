
import sys
import os
from codegen.boilerplate import *
from string import Formatter
from codegen.context import *
# name: table_rpc_events
# type: Vec<struct_rpc_events>
# init: table_rpc_events = Vec::new()
#
def fill_internal_states(definition, declaration, name, type, init, process, proto):
    assert(len(name) == len(type))
    proto_fc = proto[0].upper() + proto[1:]
    return {
        "ProtoDefinition": proto,
        # todo! field name should be configurable
        # todo! multiple type should be supported
        "ProtoGetters": f"""
fn {proto}_request_name_readonly(req: &{proto}::{proto_fc}Request) -> String {{
    let buf = &req.name as &[u8];
    String::from_utf8_lossy(buf).to_string().clone()
}}
        """,
        "ProtoRpcRequestType": f"{proto}::{proto_fc}Request",
        "ProtoRpcResponseType": f"{proto}::{proto_fc}Response",
        "InternalStatesDefinition": "".join(definition),
        "InternalStatesDeclaration": "".join([f"use crate::engine::{i};\n" for i in declaration]),
        "InternalStatesOnBuild": "".join([f"let mut {i};\n" if '=' in i else f"{i};\n" for i in init]),
        "InternalStatesOnRestore":"".join([f"let mut {i};\n" if '=' in i else f"{i};\n" for i in init]),
        "InternalStatesOnDecompose": "",
        "InternalStatesInConstructor": "".join([f"{i},\n" for i in name]),
        "InternalStatesInStructDefinition": "".join([f"pub(crate) {i[0]}:{i[1]},\n" for i in zip(name, type)]),
        "OnTxRpc": "".join(process),
        "OnRxRpc": r"""// todo """ 
    }

def retrieve_info(ctx: Context):
    proto = "hello"
    proto_fc = "Hello"
    info = {
        "ProtoDefinition": proto,
        # todo! field name should be configurable
        # todo! multiple type should be supported
        "ProtoGetters": f"""
        """,
        "ProtoRpcRequestType": f"{proto}::{proto_fc}Request",
        "ProtoRpcResponseType": f"{proto}::{proto_fc}Response",
        "InternalStatesDefinition": "\n".join(ctx.def_code),
        "InternalStatesDeclaration": "\n".join([f"use crate::engine::{i};" for i in ctx.gen_struct_names()]),
        "InternalStatesOnBuild": "\n".join(ctx.gen_init_localvar()),
        "InternalStatesOnRestore":"\n".join(ctx.gen_init_localvar()),
        "InternalStatesOnDecompose": "",
        "InternalStatesInConstructor": "\n".join([f"{i}," for i in ctx.gen_var_names()]),
        "InternalStatesInStructDefinition": "\n".join([f"pub(crate) {i}" for i in ctx.gen_struct_declaration()]),
        "OnTxRpc": "".join(ctx.process_code),
        "OnRxRpc": r"""// todo """ 
    }
    # for k,v in info.items():
    #     print(k)
    #     print(v)
    return info

def parse_intermediate_code(name):
    ctx = {
        "definition": [],
        "declaration": [],
        "internal": [],
        "init": [],
        "name": [],
        "type": [],
        "process": [],
    }
    print("Generating code for " + name)
    with open("./generated/" + name + ".rs") as f:
        current = "process"
        for i in f.readlines():
            if i.startswith("///@@"):
                j = i.split()
                if j[1] == "BEG_OF":
                    if j[2] == "declaration":
                        current = "declaration"
                    elif j[2] == "internal":
                        print("Warning: No Internal Should Be Generated")
                        current = "internal"
                    elif j[2] == "init":
                        current = "init"
                    elif j[2] == "process":
                        current = "process"
                    elif j[2] == "type":
                        current = "type"
                    elif j[2] == "name":
                        current = "name"
                    elif j[2] == "definition":
                        current = "definition"
                elif j[1] == "END_OF":
                    assert(j[2] == current)
                    current = "process"
            else:
                if current is not None:
                    if i.strip() != "":
                        ctx[current].append(i.strip('\n'))

    ctx = fill_internal_states(ctx["definition"], ctx["declaration"], ctx["name"], ctx["type"], ctx["init"], ctx["process"], "hello")
    
    return ctx
        

def gen_template(ctx, template_name, template_name_toml, template_name_first_cap, template_name_all_cap):
    target_dir = "./generated/{}".format(template_name)
    os.system(f"rm -rf {target_dir}")
    os.system(f"mkdir -p {target_dir}")
    os.chdir(target_dir)
    ctx["TemplateName"] = template_name
    ctx["TemplateNameFirstCap"] = template_name_first_cap
    ctx["TemplateNameAllCap"] = template_name_all_cap
    ctx["TemplateNameCap"] = template_name_first_cap
    print("Current dir: {}".format(os.getcwd()))
    with open("config.rs", "w") as f:
        f.write(config_rs.format(Include=include, **ctx))
    with open("lib.rs", "w") as f:
        f.write(lib_rs.format(Include=include, **ctx))
    with open("module.rs", "w") as f:
        f.write(module_rs.format(Include=include, **ctx))
    with open("engine.rs", "w") as f:
        #print([i[1] for i in Formatter().parse(engine_rs)  if i[1] is not None])
        f.write(engine_rs.format(Include=include, **ctx))
    with open("proto.rs", "w") as f:
        f.write(proto_rs)
    with open("Cargo.toml.api", "w") as f:
        f.write(api_toml.format(TemplateName=template_name_toml))
    with open("Cargo.toml.policy", "w") as f:
        f.write(policy_toml.format(TemplateName=template_name_toml))
    print("Template {} generated".format(template_name))

def move_template(mrpc_root, template_name, template_name_toml, template_name_first_cap):
    mrpc_api = mrpc_root + "/phoenix-api/policy/";
    os.system(f"rm -rf {mrpc_api}/{template_name_toml}")
    os.system(f"cp -r {mrpc_api}/logging {mrpc_api}/{template_name_toml}")
    os.system(f"rm {mrpc_api}/{template_name_toml}/Cargo.toml")
    os.system(f"cp ./Cargo.toml.api {mrpc_api}/{template_name_toml}/Cargo.toml")
    mrpc_plugin = mrpc_root + "/plugin/policy";
    os.system(f"rm -rf {mrpc_plugin}/{template_name_toml}")
    os.system(f"mkdir -p {mrpc_plugin}/{template_name_toml}/src")  
    os.system(f"cp ./Cargo.toml.policy {mrpc_plugin}/{template_name_toml}/Cargo.toml") 
    
    os.system(f"rustfmt --edition 2018  ./config.rs")
    os.system(f"rustfmt --edition 2018  ./lib.rs")
    os.system(f"rustfmt --edition 2018  ./module.rs")
    os.system(f"rustfmt --edition 2018  ./engine.rs")
    os.system(f"rustfmt --edition 2018  ./proto.rs")
    os.system(f"cp ./config.rs {mrpc_plugin}/{template_name_toml}/src/config.rs")
    os.system(f"cp ./lib.rs {mrpc_plugin}/{template_name_toml}/src/lib.rs")
    os.system(f"cp ./module.rs {mrpc_plugin}/{template_name_toml}/src/module.rs")
    os.system(f"cp ./engine.rs {mrpc_plugin}/{template_name_toml}/src/engine.rs") 
    os.system(f"cp ./proto.rs {mrpc_plugin}/{template_name_toml}/src/proto.rs")
    print("Template {} moved to mrpc folder".format(template_name))
    
def generate(name: str, ctx: Context):
    if name == "logging":
        template_name = "nofile_logging"
        template_name_toml = "nofile-logging"
        template_name_first_cap = "NofileLogging"
        template_name_all_cap = "NOFILE_LOGGING"
    elif name == "acl":
        template_name = "hello_acl"
        template_name_toml = "hello-acl"
        template_name_first_cap = "HelloAcl"
        template_name_all_cap = "HELLO_ACL"
    elif name == "fault":
        template_name = "fault"
        template_name_toml = "fault"
        template_name_first_cap = "Fault"
        template_name_all_cap = "FAULT"
        
    #ctx = parse_intermediate_code(name)
    ctx.explain()
    info = retrieve_info(ctx)
    gen_template(info, template_name, template_name_toml, template_name_first_cap, template_name_all_cap)
    move_template("/users/banruo/phoenix/experimental/mrpc", template_name, template_name_toml, template_name_first_cap)
    

    
    
    