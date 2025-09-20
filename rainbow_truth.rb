def rainbow(freq, i)
   red   = Math.sin(freq*i + 0) * 127 + 128
   green = Math.sin(freq*i + 2*Math::PI/3) * 127 + 128
   blue  = Math.sin(freq*i + 4*Math::PI/3) * 127 + 128
   "#%02X%02X%02X" % [ red, green, blue ]
end

  def self.println_plain(str, defaults={}, opts={}, chomped)
    opts.merge!(defaults)
    set_mode(opts[:truecolor]) unless @paint_init
    filtered = str.scan(ANSI_ESCAPE)
    filtered.each_with_index do |c, i|
      color = rainbow(opts[:freq], @os+i/opts[:spread])
      if opts[:invert] then
        print c[0], Paint.color(nil, color), c[1], "\e[49m"
      else
        print c[0], Paint.color(color), c[1], "\e[39m"
      end
    end

    if chomped == nil
      @old_os = @os
      @os = @os + filtered.length/opts[:spread]
    elsif @old_os
      @os = @old_os
      @old_os = nil
    end
  end

