# Defined interactively
function compare

    while true
        set fort (fortune)
        echo "$fort" | cowsay | ./cb -seed 1 -freq .05 -spread 3.2
        sleep 2
        echo $fort | cowsay | lolcat -S 1
    end
end

# set cb_comp ./cb -seed 1 -freq .05 -spread 3.2
# while true; set fort (fortune -l); echo "$fort" | cowsay | command $cb_comp ; sleep 2; echo $fort | cowsay | lolcat -S 1; end
