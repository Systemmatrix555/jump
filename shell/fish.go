package shell

var Fish = Shell(`# Put the line below in ~/.config/fish/config.fish:
#
#   status --is-interactive; and . (jump shell | psub)
#
# The following lines are autogenerated:

function __jump_add --on-variable PWD
  status --is-command-substitution; and return
  jump chdir
end

function j
  set -l dir (jump cd $argv)
  test -d "$dir"; and cd "$dir"
end
`)
