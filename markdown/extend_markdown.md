> copy form: https://markdown.com.cn/
# Markdown 扩展语法

## Markdown 表格
要添加表，请使用三个或多个连字符（---）创建每列的标题，并使用管道（|）分隔每列。您可以选择在表的任一端添加管道。
```
| Syntax      | Description |
| ----------- | ----------- |
| Header      | Title       |
| Paragraph   | Text        |
```
呈现的输出如下所示：
| Syntax      | Description |
| ----------- | ----------- |
| Header      | Title       |
| Paragraph   | Text        |

单元格宽度可以变化，如下所示。呈现的输出将看起来相同。
```
| Syntax | Description |
| --- | ----------- |
| Header | Title |
| Paragraph | Text |
```
**Tip**: 使用连字符和管道创建表可能很麻烦。为了加快该过程，请尝试使用Markdown Tables Generator。使用图形界面构建表，然后将生成的Markdown格式的文本复制到文件中。

### 对齐

您可以通过在标题行中的连字符的左侧，右侧或两侧添加冒号（`:`），将列中的文本对齐到左侧，右侧或中心。
```
| Syntax      | Description | Test Text     |
| :---        |    :----:   |          ---: |
| Header      | Title       | Here's this   |
| Paragraph   | Text        | And more      |
```
呈现的输出如下所示：

| Syntax      | Description | Test Text     |
| :---        |    :----:   |          ---: |
| Header      | Title       | Here's this   |
| Paragraph   | Text        | And more      |

### 格式化表格中的文字

您可以在表格中设置文本格式。例如，您可以添加链接，代码（仅反引号（`）中的单词或短语，而不是代码块）和强调。

您不能添加标题，块引用，列表，水平规则，图像或HTML标签。

### 在表中转义管道字符

您可以使用表格的HTML字符代码（`&#124;`）在表中显示竖线（`|`）字符。

## Markdown 围栏代码块

Markdown基本语法允许您通过将行缩进四个空格或一个制表符来创建代码块。如果发现不方便，请尝试使用受保护的代码块。根据Markdown处理器或编辑器的不同，您将在代码块之前和之后的行上使用三个反引号（ ``` ）或三个波浪号（`~~~`）。
~~~
```
{
  "firstName": "John",
  "lastName": "Smith",
  "age": 25
}
```
~~~
呈现的输出如下所示：
~~~
{
  "firstName": "John",
  "lastName": "Smith",
  "age": 25
}
~~~
Tip:要在代码块中显示反引号？请参阅了解如何转义它们。

### 语法高亮

许多Markdown处理器都支持受围栏代码块的语法突出显示。使用此功能，您可以为编写代码的任何语言添加颜色突出显示。要添加语法突出显示，请在受防护的代码块之前的反引号旁边指定一种语言。

~~~
```json
{
  "firstName": "John",
  "lastName": "Smith",
  "age": 25
}
```
~~~

呈现的输出如下所示：

```json
{
  "firstName": "John",
  "lastName": "Smith",
  "age": 25
}
```

## Markdown 脚注

脚注使您可以添加注释和参考，而不会使文档正文混乱。当您创建脚注时，带有脚注的上标数字会出现在您添加脚注参考的位置。读者可以单击链接以跳至页面底部的脚注内容。

要创建脚注参考，请在方括号（`[^1]`）内添加插入符号和标识符。标识符可以是数字或单词，但不能包含空格或制表符。标识符仅将脚注参考与脚注本身相关联-在输出中，脚注按顺序编号。

在括号内使用另一个插入符号和数字添加脚注，并用冒号和文本（`[^1]: My footnote.`）。您不必在文档末尾添加脚注。您可以将它们放在除**列表**，**块引号**和**表**之类的其他元素之外的任何位置。
```
Here's a simple footnote,[^1] and here's a longer one.[^bignote]

[^1]: This is the first footnote.

[^bignote]: Here's one with multiple paragraphs and code.

    Indent paragraphs to include them in the footnote.

    `{ my code }`

    Add as many paragraphs as you like.
```
呈现的输出如下所示：

Here's a simple footnote,[^1] and here's a longer one.[^bignote]

[^1]: This is the first footnote.

[^bignote]: Here's one with multiple paragraphs and code.

    Indent paragraphs to include them in the footnote.

    `{ my code }`

    Add as many paragraphs as you like.

## Markdown 标题编号 {#heading-ids}

许多Markdown处理器支持标题的自定义ID - 一些Markdown处理器会自动添加它们。添加自定义ID允许您直接链接到标题并使用CSS对其进行修改。要添加自定义标题ID，请在与标题相同的行上用大括号括起该自定义ID。
```
### My Great Heading {#custom-id}
```
HTML看起来像这样：
```
<h3 id="custom-id">My Great Heading</h3>
```
### 链接到标题ID

通过创建带有数字符号（`#`）和自定义标题ID的`[标准链接](/basic-syntax/links.html)`，可以链接到文件中具有自定义ID的标题。

|Markdown|	HTML|	预览效果|
|---|---|---|
|`[Heading IDs](#heading-ids)`|	`<a href="#heading-ids">Heading IDs</a>`|	[Heading IDs](#heading-ids)|

其他网站可以通过将自定义标题ID添加到网页的完整URL（例如`[Heading IDs](https://markdown.com.cn/extended-syntax/heading-ids.html#headid)`）来链接到标题。

## Markdown 定义列表

一些Markdown处理器允许您创建术语及其对应定义的定义列表。要创建定义列表，请在第一行上键入术语。在下一行，键入一个冒号，后跟一个空格和定义。
```
First Term
: This is the definition of the first term.

Second Term
: This is one definition of the second term.
: This is another definition of the second term.
```
HTML看起来像这样：

```html
<dl>
  <dt>First Term</dt>
  <dd>This is the definition of the first term.</dd>
  <dt>Second Term</dt>
  <dd>This is one definition of the second term. </dd>
  <dd>This is another definition of the second term.</dd>
</dl>
```

呈现的输出如下所示：

First Term
: This is the definition of the first term.

Second Term
: This is one definition of the second term.
: This is another definition of the second term.

## markdown 删除线

您可以通过在单词中心放置一条水平线来删除单词。结果看起来像这样。此功能使您可以指示某些单词是一个错误，要从文档中删除。若要删除单词，请在单词前后使用两个波浪号`~~`。
```
~~世界是平坦的。~~ 我们现在知道世界是圆的。
```
呈现的输出如下所示：

~~世界是平坦的。~~ 我们现在知道世界是圆的。

## Markdown 任务列表语法

任务列表使您可以创建带有复选框的项目列表。在支持任务列表的Markdown应用程序中，复选框将显示在内容旁边。要创建任务列表，请在任务列表项之前添加破折号`-`和方括号`[ ]`，并在`[ ]`前面加上空格。要选择一个复选框，请在方括号`[x]`之间添加 `x` 。
```
- [x] Write the press release
- [ ] Update the website
- [ ] Contact the media
```
呈现的输出如下所示：

- [x] Write the press release
- [ ] Update the website
- [ ] Contact the media

## Markdown 使用 Emoji 表情

有两种方法可以将表情符号添加到Markdown文件中：将表情符号复制并粘贴到Markdown格式的文本中，或者键入emoji shortcodes。

### 复制和粘贴表情符号

在大多数情况下，您可以简单地从[Emojipedia (opens new window)](https://emojipedia.org/ "表情网站")等来源复制表情符号并将其粘贴到文档中。许多Markdown应用程序会自动以Markdown格式的文本显示表情符号。从Markdown应用程序导出的HTML和PDF文件应显示表情符号。

Tip: 如果您使用的是静态网站生成器，请确保将HTML页面编码为UTF-8。.

### 使用表情符号简码

一些Markdown应用程序允许您通过键入表情符号短代码来插入表情符号。这些以冒号开头和结尾，并包含表情符号的名称。
```
去露营了！ :tent: 很快回来。

真好笑！ :joy:
```
呈现的输出如下所示：

去露营了！⛺很快回来。

真好笑！😂

Note: 注意：您可以使用此[表情符号简码列表](https://gist.github.com/rxaviers/7360908)，但请记住，表情符号简码因应用程序而异。有关更多信息，请参阅Markdown应用程序的文档。

## 自动网址链接

许多Markdown处理器会自动将URL转换为链接。这意味着如果您输入http://www.example.com，即使您未使用方括号，您的Markdown处理器也会自动将其转换为链接。

http://www.example.com

呈现的输出如下所示：

http://www.example.com

### 禁用自动URL链接

如果您不希望自动链接URL，则可以通过将URL表示为带反引号的代码来删除该链接。
```
`http://www.example.com`
```

呈现的输出如下所示：

`http://www.example.com`