# Changelog

## [O.4.O] - 2020.05.15
## Added
- add a Funcmap function that generates ready to use funcmap to use in text/template.

## [0.3.2] - 2020-02-22
### Modified
- Migrate to github.com/pirmd/text v0.5.1 and github.com/pirmd/verify v0.4.1

## [0.3.1] - 2019-11-11
### Removed
- separate style from rest of github.com/pirmd/cli repository
- style/text becomes a standalone package

## [0.3.0] - 2019-11-10

## [0.2.0] - 2019-08-11
### Added
- add new styling functions: NoLeadingSpace, NoTrailingSpace and TrimSpace
- add support to format tables, nested lists and tabs
- add preliminary support to mdoc format
### Modified
- rename NewLine to Line
- give markdown some love, alig nas far as possible plain text format with
  markdown
- clean api and (hopefully) propose better names (style.Func -> style.FormatFn,
  style.New -> style.Chainf)

## [0.1.0] - 2019-05-11
### Added
- manipulation of often used text styling for (color)term, manpage idiom
  and (very)light markdown
