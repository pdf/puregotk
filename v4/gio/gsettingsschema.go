// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// The #GSettingsSchemaSource and #GSettingsSchema APIs provide a
// mechanism for advanced control over the loading of schemas and a
// mechanism for introspecting their content.
//
// Plugin loading systems that wish to provide plugins a way to access
// settings face the problem of how to make the schemas for these
// settings visible to GSettings.  Typically, a plugin will want to ship
// the schema along with itself and it won't be installed into the
// standard system directories for schemas.
//
// #GSettingsSchemaSource provides a mechanism for dealing with this by
// allowing the creation of a new 'schema source' from which schemas can
// be acquired.  This schema source can then become part of the metadata
// associated with the plugin and queried whenever the plugin requires
// access to some settings.
//
// Consider the following example:
//
// |[&lt;!-- language="C" --&gt;
// typedef struct
//
//	{
//	   ...
//	   GSettingsSchemaSource *schema_source;
//	   ...
//	} Plugin;
//
// Plugin *
// initialise_plugin (const gchar *dir)
//
//	{
//	  Plugin *plugin;
//
//	  ...
//
//	  plugin-&gt;schema_source =
//	    g_settings_schema_source_new_from_directory (dir,
//	      g_settings_schema_source_get_default (), FALSE, NULL);
//
//	  ...
//
//	  return plugin;
//	}
//
// ...
//
// GSettings *
// plugin_get_settings (Plugin      *plugin,
//
//	const gchar *schema_id)
//
//	{
//	  GSettingsSchema *schema;
//
//	  if (schema_id == NULL)
//	    schema_id = plugin-&gt;identifier;
//
//	  schema = g_settings_schema_source_lookup (plugin-&gt;schema_source,
//	                                            schema_id, FALSE);
//
//	  if (schema == NULL)
//	    {
//	      ... disable the plugin or abort, etc ...
//	    }
//
//	  return g_settings_new_full (schema, NULL, NULL);
//	}
//
// ]|
//
// The code above shows how hooks should be added to the code that
// initialises (or enables) the plugin to create the schema source and
// how an API can be added to the plugin system to provide a convenient
// way for the plugin to access its settings, using the schemas that it
// ships.
//
// From the standpoint of the plugin, it would need to ensure that it
// ships a gschemas.compiled file as part of itself, and then simply do
// the following:
//
// |[&lt;!-- language="C" --&gt;
//
//	{
//	  GSettings *settings;
//	  gint some_value;
//
//	  settings = plugin_get_settings (self, NULL);
//	  some_value = g_settings_get_int (settings, "some-value");
//	  ...
//	}
//
// ]|
//
// It's also possible that the plugin system expects the schema source
// files (ie: .gschema.xml files) instead of a gschemas.compiled file.
// In that case, the plugin loading system must compile the schemas for
// itself before attempting to create the settings source.
type SettingsSchema struct {
}

// #GSettingsSchemaKey is an opaque data structure and can only be accessed
// using the following functions.
type SettingsSchemaKey struct {
}

// This is an opaque structure type.  You may not access it directly.
type SettingsSchemaSource struct {
}

var xSettingsSchemaSourceGetDefault func() *SettingsSchemaSource

// Gets the default system schema source.
//
// This function is not required for normal uses of #GSettings but it
// may be useful to authors of plugin management systems or to those who
// want to introspect the content of schemas.
//
// If no schemas are installed, %NULL will be returned.
//
// The returned source may actually consist of multiple schema sources
// from different directories, depending on which directories were given
// in `XDG_DATA_DIRS` and `GSETTINGS_SCHEMA_DIR`. For this reason, all
// lookups performed against the default source should probably be done
// recursively.
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {

	cret := xSettingsSchemaSourceGetDefault()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xSettingsSchemaSourceGetDefault, lib, "g_settings_schema_source_get_default")

}
