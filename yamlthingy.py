"""
YAML Templating Thingy
-- Tim Dousset 2015-04-28
-- Shaun Smekel 2015-05-14
--- Added dynamic environment file imports 
"""

import re
import os
import sys
import argparse
import yaml
import time
import datetime

"""
Handles splicing YAML files
"""
class YAMLThingy(object):
  def __init__(self, templatefile, marker, environmentPath):

    # Create timestamp to inject for template identifier
    ts = time.time()
    timestamp = datetime.datetime.fromtimestamp(ts).strftime('%Y-%m-%d %H:%M:%S')

    outputfile = '%s.tmp' % templatefile
    self._tf = file(templatefile, 'r')
    try:
      yaml.load(self._tf)
    except yaml.YAMLError, exc:
      print "Could not parse templatefile YAML"
      if hasattr(exc, 'problem_mark'):
        mark = exc.problem_mark
        print "Error position: (%s:%s)" % (mark.line+1, mark.column+1)
      sys.exit(1)
    print "templatefile %s is validated." % (templatefile, )

    # Init the output file
    self._of = file(outputfile, 'w')

    # Verify that the pattern occurs.
    found = False
    self._tf.seek(0)
    for line in self._tf.readlines():
      self._of.write(line)

      # Check for any environmental files
      envMatch = re.match('.*%s.*' % (marker, ), line)
      if envMatch is not None:
        found = True

        print "Match is %s" % envMatch.group(1)

        if envMatch.group(1) == 'timestamp':
          
          print "Attaching timestamp to property"

          # Match indents for valid YAML file
          indent = ''
          indentmatch = re.match('(\s+)', line)
          
          if indentmatch is not None:
            indent = indentmatch.group(0)
          
          newline = '%s  "%s"\n' % (indent, timestamp)
          self._of.write(newline)

        else:
          # Check for environment file specified
          environmentFilename =  '%s%s.yaml' % (environmentPath, envMatch.group(1))
          try:
            self._cf = file(environmentFilename, 'r')
          except IOError, exc: 
            print "Could not find environment file %s" % environmentFilename
            continue

          try:
            yaml.load(self._cf)
          except yaml.YAMLError, exc:
            print "Could not parse environment file %s" % environmentFilename
            if hasattr(exc, 'problem_mark'):
              mark = exc.problem_mark
              print "Error position: (%s:%s)" % (mark.line + 1, mark.column + 1)
            sys.exit(1)
          print "environment file %s is validated." % environmentFilename

          # Match indents for valid YAML file
          indent = ''
          indentmatch = re.match('(\s+)', line)
          
          if indentmatch is not None:
            indent = indentmatch.group(0)
          self._cf.seek(0)
          

          for contentline in self._cf.readlines():
            newline = '%s  %s' % (indent, contentline)
            self._of.write(newline)
    
    self._of.close()

    # Move the temp output file to override the original template
    os.rename(outputfile, templatefile)

    # if not found:
      # print "Did not find the marker '%s' within the template file!" % (marker, )

if __name__ == "__main__":
  parser = argparse.ArgumentParser()
  parser.add_argument('--templatefile', dest='templatefile', default=None,
                      help='Template YAML file.')
  parser.add_argument('--marker', dest='marker', default='([a-zA-Z0-9-_]+):\\s#\\1#',
                      help='Marker identify where to inject environment file contents. \
                      The default ([a-zA-Z0-9-_]+):\\s#\\1# will match string: #string# \
                      and will import the contents of string.yaml from the environment directory.')
  parser.add_argument('--envdir', dest='envdir', default=None,
                      help='File containing YAML data to inject.')
  # parser.add_argument('--outputfile', dest='outputfile', default=None,
  #                     help='File to write final product out to.')
  args = parser.parse_args()

  if args.templatefile is not None:
    if not os.path.isfile(args.templatefile):
      print "Could not find the templatefile you specified."
      print args.templatefile
      sys.exit(1)
  else:
    print "You must supply a templatefile (--templatefile)"
    parser.print_help()
    sys.exit(1)

  if args.marker is None or len(args.marker) == 0:
    print "You must supply a marker included in the template to show where contentfile should be injected."
    sys.exit(1)

  if args.envdir is not None:
    if not os.path.isdir(args.envdir):
      print "Could not find the envdir you specified."
      print args.envdir
      sys.exit(1)
  else:
    print "You must supply an envdir (--envdir)"
    parser.print_help()
    sys.exit(1)

  # if args.outputfile is None:
  #   print "You must supply a outputfile (--outputfile)"
  #   parser.print_help()
  #   sys.exit(1)

  try:
    # thingy = YAMLThingy(args.templatefile, args.marker, args.envdir, args.outputfile)
    thingy = YAMLThingy(args.templatefile, args.marker, args.envdir)
  except KeyboardInterrupt:
    print 'Interrupted. Bailing...'
    sys.exit(1)
