> copy form: https://markdown.com.cn/

# Markdown 基本语法

## Markdown 标题语法

要创建标题，请在单词或短语前面添加井号 (`#`) 。`#` 的数量代表了标题的级别。例如，添加三个 `#` 表示创建一个三级标题 (`<h3>`) (例如：`### My Header`)。

| Markdown语法                 | HTML                           | 预览效果                |
| ------------------------------ | -------------------------------- | ------------------------- |
| `# Heading level 1`      | `Heading level 1` | # Heading level 1      |
| `## Heading level 2`     | `Heading level 2` | ## Heading level 2     |
| `### Heading level 3`    | `Heading level 3` | ### Heading level 3    |
| `#### Heading level 4`   | `Heading level 4` | #### Heading level 4   |
| `##### Heading level 5`  | `Heading level 5` | ##### Heading level 5  |
| `###### Heading level 6` | `Heading level 6` | ###### Heading level 6 |

### 可选语法

还可以在文本下方添加任意数量的 == 号来标识一级标题，或者 -- 号来标识二级标题。

| Markdown语法                              | HTML                           | 预览效果            |
| ------------------------------------------- | -------------------------------- | --------------------- |
| `Heading level 1===============` | `Heading level 1` | # Heading level 1  |
| `Heading level 2---------------` | `Heading level 2` | ## Heading level 2 |

### 最佳实践

不同的 Markdown 应用程序处理 `#` 和标题之间的空格方式并不一致。为了兼容考虑，请用一个空格在 `#` 和标题之间进行分隔。

| ✅  Do this              | ❌  Don't do this       |
| -------------------------- | ------------------------- |
| `# Here's a Heading` | `#Here's a Heading` |

## Markdown 段落

要创建段落，请使用空白行将一行或多行文本进行分隔。

| Markdown语法                  | HTML                | 预览效果     |
| ------------------------------ | ----------------------- | ------------------ |
| `I really like using Markdown.I think I'll use it to format all of my documents from now on.` | `I really like using Markdown.I think I'll use it to format all of my documents from now on.` | I really like using Markdown.  <br> I think I'll use it to format all of my documents from now on. |

### 段落（Paragraph）用法的最佳实

不要用空格（spaces）或制表符（ tabs）缩进段落。

| ✅  Do this    | ❌  Don't do this    |
| ------------------------ | ---------------------------------- |
| `Don't put tabs or spaces in front of your paragraphs.Keep lines left-aligned like this.` | `    This can result in unexpected formatting problems.  Don't add tabs or spaces in front of paragraphs.` |

## Markdown 换行语法

在一行的末尾添加两个或多个空格，然后按回车键,即可创建一个换行(`<br>`)。

| Markdown语法         | HTML           | 预览效果        |
| --------------------- | ------------------------------------- | -------------------------------------- |
| `This is the first line.  And this is the second line.` | `This is the first line.And this is the second line.` | This is the first line.And this is the second line. |

### 换行（Line Break）用法的最佳实践

几乎每个 Markdown 应用程序都支持两个或多个空格进行换行，称为 `结尾空格（trailing whitespace)` 的方式，但这是有争议的，因为很难在编辑器中直接看到空格，并且很多人在每个句子后面都会有意或无意地添加两个空格。由于这个原因，你可能要使用除结尾空格以外的其它方式来换行。幸运的是，几乎每个 Markdown 应用程序都支持另一种换行方式：HTML 的 `<br>` 标签。

为了兼容性，请在行尾添加“结尾空格”或 HTML 的 `<br>` 标签来实现换行。

还有两种其他方式我并不推荐使用。CommonMark 和其它几种轻量级标记语言支持在行尾添加反斜杠 (`\`) 的方式实现换行，但是并非所有 Markdown 应用程序都支持此种方式，因此从兼容性的角度来看，不推荐使用。并且至少有两种轻量级标记语言支持无须在行尾添加任何内容，只须键入回车键（`return`）即可实现换行。

| ✅  Do this  | ❌  Don't do this         |
| ------------------------- | -------------------------------------------------- |
| `First line with two spaces after.  And the next line.First line with the HTML tag after.And the next line.` | `First line with a backslash after.\And the next line.First line with nothing after.` |

## Markdown 强调语法

通过将文本设置为粗体或斜体来强调其重要性。

### 粗体（Bold）

要加粗文本，请在单词或短语的前后各添加两个星号（asterisks）或下划线（underscores）。如需加粗一个单词或短语的中间部分用以表示强调的话，请在要加粗部分的两侧各添加两个星号（asterisks）。

| Markdown语法                     | HTML                                          | 预览效果                             |
| ---------------------------------- | ----------------------------------------------- | -------------------------------------- |
| `I just love **bold text**.` | `I just love bold text.` | I just love ​**bold text**​. |
| `I just love __bold text__.` | `I just love bold text.` | I just love ​**bold text**​. |
| `Love**is**bold`             | `Loveisbold`             | Love**is**bold                 |

#### 粗体（Bold）用法最佳实践

Markdown 应用程序在如何处理单词或短语中间的下划线上并不一致。为兼容考虑，在单词或短语中间部分加粗的话，请使用星号（asterisks）。

| ✅  Do this          | ❌  Don't do this    |
| ---------------------- | ---------------------- |
| `Love**is**bold` | `Love__is__bold` |

### 斜体（Italic）

要用斜体显示文本，请在单词或短语前后添加一个星号（asterisk）或下划线（underscore）。要斜体突出单词的中间部分，请在字母前后各添加一个星号，中间不要带空格。

| Markdown语法                               | HTML                                              | 预览效果                                    |
| -------------------------------------------- | --------------------------------------------------- | --------------------------------------------- |
| `Italicized text is the *cat's meow*.` | `Italicized text is the cat's meow.` | Italicized text is the ​*cat’s meow*​. |
| `Italicized text is the _cat's meow_.` | `Italicized text is the cat's meow.` | Italicized text is the ​*cat’s meow*​. |
| `A*cat*meow`                           | `Acatmeow`                           | A*cat*meow                                |

#### 斜体（Italic）用法的最佳实践

要同时用粗体和斜体突出显示文本，请在单词或短语的前后各添加三个星号或下划线。要加粗并用斜体显示单词或短语的中间部分，请在要突出显示的部分前后各添加三个星号，中间不要带空格。

| ✅  Do this      | ❌  Don't do this |
| ------------------ | ------------------- |
| `A*cat*meow` | `A_cat_meow`  |

### 粗体（Bold）和斜体（Italic）

要同时用粗体和斜体突出显示文本，请在单词或短语的前后各添加三个星号或下划线。要加粗并用斜体显示单词或短语的中间部分，请在要突出显示的部分前后各添加三个星号，中间不要带空格。

| Markdown语法                                  | HTML                                                              | 预览效果                                         |
| ----------------------------------------------- | ------------------------------------------------------------------- | -------------------------------------------------- |
| `This text is ***really important***.`    | `This text is really important.`    | This text is ​***really important***​. |
| `This text is ___really important___.`    | `This text is really important.`    | This text is ​***really important***​. |
| `This text is __*really important*__.`    | `This text is really important.`    | This text is ​***really important***​. |
| `This text is **_really important_**.`    | `This text is really important.`    | This text is ​***really important***​. |
| `This is really***very***important text.` | `This is reallyveryimportant text.` | This is really***very***important text.  |

#### 粗体（Bold）和斜体（Italic）用法的最佳实践

Markdown 应用程序在处理单词或短语中间添加的下划线上并不一致。为了实现兼容性，请使用星号将单词或短语的中间部分加粗并以斜体显示，以示重要。

| ✅  Do this                                   | ❌  Don't do this                             |
| ----------------------------------------------- | ----------------------------------------------- |
| `This is really***very***important text.` | `This is really___very___important text.` |

## Markdown 引用语法

要创建块引用，请在段落前添加一个 `>` 符号。

```text
> Dorothy followed her through many of the beautiful rooms in her castle.
```

渲染效果如下所示：

> Dorothy followed her through many of the beautiful rooms in her castle.

#### 多个段落的块引用

块引用可以包含多个段落。为段落之间的空白行添加一个 `>` 符号。

```text
> Dorothy followed her through many of the beautiful rooms in her castle.
>
> The Witch bade her clean the pots and kettles and sweep the floor and keep the fire fed with wood.
```

渲染效果如下：

> Dorothy followed her through many of the beautiful rooms in her castle.
> 
> The Witch bade her clean the pots and kettles and sweep the floor and keep the fire fed with wood.

#### 嵌套块引用

块引用可以嵌套。在要嵌套的段落前添加一个 `>>` 符号。

```text
> Dorothy followed her through many of the beautiful rooms in her castle.
>
>> The Witch bade her clean the pots and kettles and sweep the floor and keep the fire fed with wood.
```

渲染效果如下：

> Dorothy followed her through many of the beautiful rooms in her castle.
> 
> > The Witch bade her clean the pots and kettles and sweep the floor and keep the fire fed with wood.

#### 带有其它元素的块引用

块引用可以包含其他 Markdown 格式的元素。并非所有元素都可以使用，你需要进行实验以查看哪些元素有效。

```text
> #### The quarterly results look great!
>
> - Revenue was off the chart.
> - Profits were higher than ever.
>
>  *Everything* is going according to **plan**.
```

渲染效果如下：

> #### The quarterly results look great!
> 
> * Revenue was off the chart.
> * Profits were higher than ever.
> 
> *Everything* is going according to ​**plan**​.

## Markdown 列表语法

可以将多个条目组织成有序或无序列表。

### 有序列表

要创建有序列表，请在每个列表项前添加数字并紧跟一个英文句点。数字不必按数学顺序排列，但是列表应当以数字 1 起始。

| Markdown语法            | HTML                    | 预览效果              |
| ------------------------------- | -------------------------------------------------- | ------------------------------------------ |
| `1. First item2. Second item3. Third item4. Fourth item`                                                           | `First itemSecond itemThird itemFourth item`                                                                               | 1. First item1. Second item1. Third item1. Fourth item                                 |
| `1. First item1. Second item1. Third item1. Fourth item`                                                           | `First itemSecond itemThird itemFourth item`                                                                               | 1. First item1. Second item1. Third item1. Fourth item                                 |
| `1. First item8. Second item3. Third item5. Fourth item`                                                           | `First itemSecond itemThird itemFourth item`                                                                               | 1. First item1. Second item1. Third item1. Fourth item                                 |
| `1. First item2. Second item3. Third item    1. Indented item    2. Indented item4. Fourth item` | `First itemSecond itemThird itemIndented itemIndented itemFourth item` | 1. First item1. Second item1. Third item1. Indented item1. Indented item1. Fourth item |

#### 有序列表最佳实践

CommonMark and a few other lightweight markup languages let you use a parenthesis (`)`) as a delimiter (e.g., `1) First item`), but not all Markdown applications support this, so it isn’t a great option from a compatibility perspective. For compatibility, use periods only.

| ✅  Do this                            | ❌  Don't do this                      |
| ---------------------------------------- | ---------------------------------------- |
| `1. First item2. Second item` | `1) First item2) Second item` |

### 无序列表

要创建无序列表，请在每个列表项前面添加破折号 (-)、星号 (\*) 或加号 (+) 。缩进一个或多个列表项可创建嵌套列表。

| Markdown语法            | HTML                 | 预览效果                        |
| -------- | --------------------- | -------------------------------------- |
| `- First item- Second item- Third item- Fourth item`                                                         | `First itemSecond itemThird itemFourth item`                                                                               | * First item* Second item* Third item* Fourth item                                 |
| `* First item* Second item* Third item* Fourth item`                                                         | `First itemSecond itemThird itemFourth item`                                                                               | * First item* Second item* Third item* Fourth item                                 |
| `+ First item+ Second item+ Third item+ Fourth item`                                                         | `First itemSecond itemThird itemFourth item`                                                                               | * First item* Second item* Third item* Fourth item                                 |
| `- First item- Second item- Third item    - Indented item    - Indented item- Fourth item` | `First itemSecond itemThird itemIndented itemIndented itemFourth item` | * First item* Second item* Third item* Indented item* Indented item* Fourth item |


#### 无序列表最佳实践

Markdown applications don’t agree on how to handle different delimiters in the same list. For compatibility, don't mix and match delimiters in the same list — pick one and stick with it.

| ✅  Do this                                                             | ❌  Don't do this                                                       |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `- First item- Second item- Third item- Fourth item` | `+ First item* Second item- Third item+ Fourth item` |

### 在列表中嵌套其他元素

要在保留列表连续性的同时在列表中添加另一种元素，请将该元素缩进四个空格或一个制表符，如下例所示：

#### 段落

```text
*   This is the first list item.
*   Here's the second list item.

    I need to add another paragraph below the second list item.

*   And here's the third list item.
```

渲染效果如下：

* This is the first list item.
* Here's the second list item.
  I need to add another paragraph below the second list item.
* And here's the third list item.

#### 引用块

```text
*   This is the first list item.
*   Here's the second list item.

    > A blockquote would look great below the second list item.

*   And here's the third list item.
```

渲染效果如下：

* This is the first list item.
* Here's the second list item.
  > A blockquote would look great below the second list item.
* And here's the third list item.

#### 代码块

代码块通常采用四个空格或一个制表符缩进。当它们被放在列表中时，请将它们缩进八个空格或两个制表符。

```text
1.  Open the file.
2.  Find the following code block on line 21:

        &lt;html>
          &lt;head>
            &lt;title>Test&lt;/title>
          &lt;/head>

3.  Update the title to match the name of your website.
```

渲染效果如下：

1. Open the file.
2. Find the following code block on line 21:
   ```text
   &lt;html>
     &lt;head>
       &lt;title>Test&lt;/title>
     &lt;/head>
   ```
3. Update the title to match the name of your website.

#### 图片

```text
1.  Open the file containing the Linux mascot.
2.  Marvel at its beauty.

    ![Tux, the Linux mascot](/assets/images/tux.png)

3.  Close the file.
```

渲染效果如下：

1. Open the file containing the Linux mascot.
2. Marvel at its beauty.
   ![Tux, the Linux mascot](imagepath)
3. Close the file.

#### 列表

You can nest an unordered list in an ordered list, or vice versa.

```text
1. First item
2. Second item
3. Third item
    - Indented item
    - Indented item
4. Fourth item
```

渲染效果如下：

1. First item
2. Second item
3. Third item
   * Indented item
   * Indented item
4. Fourth item

## Markdown 代码语法

要将单词或短语表示为代码，请将其包裹在反引号 (`) 中。

|Markdown语法	|HTML	|预览效果|
|----|----|---|
|At the command prompt, type `nano`.	|At the command prompt, type <code>nano</code>.	|At the command prompt, type nano.|

### 转义反引号
如果你要表示为代码的单词或短语中包含一个或多个反引号，则可以通过将单词或短语包裹在双反引号(``)中。

``Use `code` in your Markdown file.``

### 代码块
要创建代码块，请将代码块的每一行缩进至少四个空格或一个制表符。
```
    &lt;html>
      &lt;head>
      &lt;/head>
    &lt;/html>
```
渲染效果如下：

&lt;html>
  &lt;head>
  &lt;/head>
&lt;/html>

Note: 要创建不用缩进的代码块，请使用 围栏式代码块（fenced code blocks）.

## Markdown 分隔线语法

要创建分隔线，请在单独一行上使用三个或多个星号 (***)、破折号 (---) 或下划线 (___) ，并且不能包含其他内容。
```
***

---

_________________
```
以上三个分隔线的渲染效果看起来都一样：

### 分隔线（Horizontal Rule）用法最佳实践

为了兼容性，请在分隔线的前后均添加空白行。

Try to put a blank line before...

---

...and after a horizontal rule.	Without blank lines, this would be a heading.

## Markdown 链接语法
链接文本放在中括号内，链接地址放在后面的括号中，链接title可选。

超链接Markdown语法代码：`[超链接显示名](超链接地址 "超链接title")`

对应的HTML代码：`<a href="超链接地址" title="超链接title">超链接显示名</a>`

这是一个链接 `[Markdown语法](https://markdown.com.cn)`。

渲染效果如：[这是一个链接](https//tets.com "这是一个链接titile")

### 给链接增加 Title

链接title是当鼠标悬停在链接上时会出现的文字，这个title是可选的，它放在圆括号中链接地址后面，跟链接地址之间以空格分隔。

 `[Markdown语法](https://markdown.com.cn "最好的markdown教程")`。

渲染效果如：[这是一个链接](https//tets.com "这是一个链接titile")

### 网址和Email地址

使用尖括号可以很方便地把URL或者email地址变成可点击的链接。

`<https://markdown.com.cn>`
`<fake@example.com>`

渲染效果如：<fake@example.com>

### 带格式化的链接

强调链接, 在链接语法前后增加星号。
要将链接表示为代码，请在方括号中添加反引号。
```
I love supporting the **[EFF](https://eff.org)**.
This is the *[Markdown Guide](https://www.markdownguide.org)*.
See the section on [`code`](#code).
```
渲染效果如下：

- I love supporting the **[EFF](https://eff.org)**.
- This is the *[Markdown Guide](https://www.markdownguide.org)*.
- See the section on [`code`](#code).

### 引用类型链接
引用样式链接是一种特殊的链接，它使URL在Markdown中更易于显示和阅读。

参考样式链接分为两部分：**与文本保持内联的部分**以及**存储在文件中其他位置的部分**，以使文本易于阅读。

#### 链接的第一部分格式

引用类型的链接的第一部分使用两组括号进行格式设置。第一组方括号包围应显示为链接的文本。第二组括号显示了一个标签，该标签用于指向您存储在文档其他位置的链接。

尽管不是必需的，可以在第一组和第二组括号之间包含一个空格。第二组括号中的标签不区分大小写，可以包含字母，数字，空格或标点符号。

以下示例格式对于链接的第一部分效果相同：
- `[hobbit-hole][1]`
- `[hobbit-hole] [1]`

#### 链接的第二部分格式

引用类型链接的第二部分使用以下属性设置格式：
1. 放在括号中的标签，其后紧跟一个冒号和至少一个空格（例如[label]:）。
2. 链接的URL，可以选择将其括在尖括号中。
3. 链接的可选标题，可以将其括在双引号，单引号或括号中。

以下示例格式对于链接的第二部分效果相同：
- `[1]: https://en.wikipedia.org/wiki/Hobbit#Lifestyle`
- `[1]: https://en.wikipedia.org/wiki/Hobbit#Lifestyle "Hobbit lifestyles"`
- `[1]: https://en.wikipedia.org/wiki/Hobbit#Lifestyle 'Hobbit lifestyles'`
- `[1]: https://en.wikipedia.org/wiki/Hobbit#Lifestyle (Hobbit lifestyles)`
- `[1]: <https://en.wikipedia.org/wiki/Hobbit#Lifestyle> "Hobbit lifestyles"`
- `[1]: <https://en.wikipedia.org/wiki/Hobbit#Lifestyle> 'Hobbit lifestyles'`
- `[1]: <https://en.wikipedia.org/wiki/Hobbit#Lifestyle> (Hobbit lifestyles)`

可以将链接的第二部分放在Markdown文档中的任何位置。有些人将它们放在出现的段落之后，有些人则将它们放在文档的末尾（例如尾注或脚注）。

### 链接最佳实践
不同的 Markdown 应用程序处理URL中间的空格方式不一样。为了兼容性，请尽量使用%20代替空格。
`[link](https://www.example.com/my%20great%20page)`

## Markdown 图片语法

要添加图像，请使用感叹号 (!), 然后在方括号增加替代文本，图片链接放在圆括号里，括号里的链接后可以增加一个可选的图片标题文本。

插入图片Markdown语法代码：`![图片alt](图片链接 "图片title")`。

对应的HTML代码：`<img src="图片链接" alt="图片alt" title="图片title">`

渲染效果如：`![这是图片](/assets/img/philly-magic-garden.jpg "Magic Gardens")`

### 链接图片

给图片增加链接，请将图像的Markdown 括在方括号中，然后将链接添加在圆括号中。

`[![沙漠中的岩石图片](/assets/img/shiprock.jpg "Shiprock")](https://markdown.com.cn)`

渲染效果如：`[![沙漠中的岩石图片](/assets/img/shiprock.jpg "Shiprock")](https://markdown.com.cn)`


## Markdown 转义字符语法
要显示原本用于格式化 Markdown 文档的字符，请在字符前面添加反斜杠字符 \ 。

`\* Without the backslash, this would be a bullet in an unordered list.`

渲染效果如下：

\* Without the backslash, this would be a bullet in an unordered list.

### 可做转义的字符

以下列出的字符都可以通过使用反斜杠字符从而达到转义目的。

|Character|	Name|
|-----|---|
|\	|backslash|
|`	|backtick (see also escaping backticks in code)|
|*	|asterisk|
|_	|underscore|
|{ }	|curly braces|
|[ ]	|brackets|
|( )	|parentheses|
|#	|pound sign|
|+	|plus sign|
|-	|minus sign (hyphen)|
|.	|dot|
|!	|exclamation mark|
||	|pipe (see also escaping pipe in tables)|

### 特殊字符自动转义

在 HTML 文件中，有两个字符需要特殊处理： `<` 和 `&` 。 `<` 符号用于起始标签，`&` 符号则用于标记 HTML 实体，如果你只是想要使用这些符号，你必须要使用实体的形式，像是 `&lt;` 和 `&amp;`。

`&` 符号其实很容易让写作网页文件的人感到困扰，如果你要打`「AT&T」` ，你必须要写成`「AT&amp;T」` ，还得转换网址内的 `&` 符号，如果你要链接到：`http://images.google.com/images?num=30&q=larry+bird`
你必须要把网址转成：`http://images.google.com/images?num=30&amp;q=larry+bird`才能放到链接标签的 href 属性里。不用说也知道这很容易忘记，这也可能是 HTML 标准检查所检查到的错误中，数量最多的。

Markdown 允许你直接使用这些符号，它帮你自动转义字符。如果你使用 `&` 符号的作为 HTML 实体的一部分，那么它不会被转换，而在其它情况下，它则会被转换成 `&amp;`。所以你如果要在文件中插入一个著作权的符号，你可以这样写：`&copy;`Markdown 将不会对这段文字做修改，但是如果你这样写：`AT&T`，·Markdown 就会将它转为：`AT&amp;T`

类似的状况也会发生在 `<` 符号上，因为 Markdown 支持 行内 HTML ，如果你使用 `<` 符号作为 HTML 标签的分隔符，那 Markdown 也不会对它做任何转换，但是如果你是写：`4 < 5`,Markdown 将会把它转换为：`4 &lt; 5`

需要特别注意的是，在 Markdown 的块级元素和内联元素中， `<` 和 `&` 两个符号都会被自动转换成 HTML 实体，这项特性让你可以很容易地用 Markdown 写 HTML。（在 HTML 语法中，你要手动把所有的 `<` 和 `&` 都转换为 HTML 实体。）

## Markdown 内嵌 HTML 标签

对于 Markdown 涵盖范围之外的标签，都可以直接在文件里面用 HTML 本身。如需使用 HTML，不需要额外标注这是 HTML 或是 Markdown，只需 HTML 标签添加到 Markdown 文本中即可。

### 行级內联标签

HTML 的行级內联标签如 `<span>`、`<cite>`、`<del>` 不受限制，可以在 Markdown 的段落、列表或是标题里任意使用。依照个人习惯，甚至可以不用 Markdown 格式，而采用 HTML 标签来格式化。例如：如果比较喜欢 HTML 的 `<a>` 或 `<img>` 标签，可以直接使用这些标签，而不用 Markdown 提供的链接或是图片语法。当你需要更改元素的属性时（例如为文本指定颜色或更改图像的宽度），使用 HTML 标签更方便些。

HTML 行级內联标签和区块标签不同，在內联标签的范围内， Markdown 的语法是可以解析的。

`This **word** is bold. This <em>word</em> is italic.`

渲染效果如下:

This **word** is bold. This <em>word</em> is italic.

### 区块标签

区块元素──比如 `<div>`、`<table>`、`<pre>`、`<p>` 等标签，必须在前后加上空行，以便于内容区分。而且这些元素的开始与结尾标签，不可以用 `tab` 或是空白来缩进。Markdown 会自动识别这区块元素，避免在区块标签前后加上没有必要的 `<p>` 标签。

例如，在 Markdown 文件里加上一段 HTML 表格：

```
This is a regular paragraph.

<table>
    <tr>
        <td>Foo</td>
    </tr>
</table>

This is another regular paragraph.
```
渲染效果如下：

This is a regular paragraph.

<table>
    <tr>
        <td>Foo</td>
    </tr>
</table>

This is another regular paragraph.

请注意，Markdown 语法在 HTML 区块标签中将不会被进行处理。例如，你无法在 HTML 区块内使用 Markdown 形式的`*强调*`。

### HTML 用法最佳实践
出于安全原因，并非所有 Markdown 应用程序都支持在 Markdown 文档中添加 HTML。如有疑问，请查看相应 Markdown 应用程序的手册。某些应用程序只支持 HTML 标签的子集。

对于 HTML 的块级元素 `<div>`、`<table>`、`<pre>` 和 `<p>`，请在其前后使用空行（blank lines）与其它内容进行分隔。尽量不要使用制表符（tabs）或空格（spaces）对 HTML 标签做缩进，否则将影响格式。

在 HTML 块级标签内不能使用 Markdown 语法。例如 `<p>italic and **bold**</p>` 将不起作用。